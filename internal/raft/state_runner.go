package raft

import (
	"configStorage/api/raftrpc"
	"configStorage/tools/random"
	"context"
	"sync"
	"time"
)

func (rf *Raft) follower(ctx context.Context) {
	rf.logger.Printf("start state as follower")
	for {
		select {
		case <-ctx.Done():
			rf.logger.Printf("finish state as follower")
			return
		default:
			rf.loopFollower()
		}
	}
}

func (rf *Raft) candidate(ctx context.Context) {
	rf.logger.Printf("start state as candidate")
	for {
		select {
		case <-ctx.Done():
			rf.logger.Printf("finish state as candidate")
			return
		default:
			rf.loopCandidate()
		}
	}
}

func (rf *Raft) leader(ctx context.Context) {
	rf.logger.Printf("start state as leader")
	heartbeatOk := make(chan struct{})
	rentDue := make(chan struct{})

	go func() {
		// if state is a leader, then set a rent for seconds
		// when rent is due, finish the instance as a leader,
		// then start a new round of election
		for {
			select {
			case <-time.After(LeaderRentTimeout):
				close(rentDue)
				return
			case <-heartbeatOk:
				continue
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			rf.logger.Printf("finish state as leader")
			return
		case <-rentDue:
			rf.logger.Printf("rent is due as a leader")
			rf.mu.Lock()
			rf.state = Follower
			rf.stateChange <- Follower
			rf.mu.Unlock()
			return
		default:
			if rf.loopLeader() {
				heartbeatOk <- struct{}{}
			}
		}
	}
}

func (rf *Raft) loopFollower() {
	rf.mu.Lock()
	// reset info as a follower
	rf.votedFor = UnVoted
	rf.heartbeat = false
	rf.mu.Unlock()

	timeout := random.Timeout(HeartBeatTimeout, RandTimeout)
	time.Sleep(timeout)

	rf.mu.Lock()
	// if after sleep and not receive from leader
	// and haven't voted for anyone, then start a new election
	if !rf.heartbeat && rf.state == Follower &&
		(rf.votedFor == UnVoted || rf.votedFor == rf.id) {
		rf.state = Candidate
		rf.stateChange <- Candidate
	}
	rf.mu.Unlock()
}

func (rf *Raft) loopCandidate() {
	rf.mu.Lock()

	if rf.state != Candidate ||
		(rf.votedFor != rf.id && rf.votedFor != UnVoted) {
		rf.mu.Unlock()
		return
	}

	rf.votedFor = rf.id
	peersNum := cap(rf.peers)
	success := make(chan struct{})
	finish := make(chan struct{}, 1)
	rf.mu.Unlock()

	go func() {
		rf.mu.Lock()
		for i, peer := range rf.peers {
			if i != int(rf.id) {
				args := raftrpc.RequestVoteArgs{
					Term:         rf.currentTerm,
					CandidateID:  rf.id,
					LastLogIndex: 0,
					LastLogTerm:  0,
				}
				if len(rf.logs) > 0 {
					args.LastLogIndex = rf.logs[len(rf.logs)-1].Index
					args.LastLogTerm = rf.logs[len(rf.logs)-1].Term
				}
				// send rpc request for votes
				go func(peerIndex int, peer raftrpc.RaftClient, args *raftrpc.RequestVoteArgs) {
					reply, err := peer.RequestVote(context.Background(), args)
					if err != nil {
						rf.logger.Printf("fail to send rpc request to id %d: %v", peerIndex, err.Error())
						return
					}
					rf.mu.Lock()
					defer rf.mu.Unlock()
					if reply.Term > rf.currentTerm {
						rf.state = Follower
						rf.votedFor = UnVoted
						rf.stateChange <- Follower
						return
					}
					if reply.VoteGranted {
						rf.logger.Printf("receive vote from peer %d", peerIndex)
						finish <- struct{}{}
						return
					}
				}(i, peer, &args)
			}
		}
		rf.mu.Unlock()
		cnt := peersNum / 2
		tot := peersNum - 1
		once := sync.Once{}
		for {
			select {
			case <-finish:
				cnt--
				tot--
				if cnt <= 0 {
					once.Do(func() {
						close(success)
					})
				}
				if tot <= 0 {
					return
				}
			}
		}
	}()

	select {
	case <-success:
		rf.mu.Lock()
		if rf.state != Candidate {
			rf.mu.Unlock()
			return
		}
		rf.currentTerm++
		for index := 0; index < peersNum; index++ {
			rf.leaderState.matchIndex[int32(index)] = rf.currentIndex
			rf.leaderState.nextIndex[int32(index)] = 0
		}
		rf.state = Leader
		rf.stateChange <- Leader
		rf.mu.Unlock()
	case <-time.After(random.Timeout(ElectionTimeout, 0)):
		rf.mu.Lock()
		if rf.state != Candidate {
			rf.mu.Unlock()
			return
		}
		rf.votedFor = UnVoted
		rf.mu.Unlock()
		// sleep and wait for another election
		timeout := random.Timeout(0, RandTimeout)
		time.Sleep(timeout)
	}
}

// loopLeader main func when state is leader
// return true if successfully contact all the other peers
func (rf *Raft) loopLeader() bool {
	rf.mu.Lock()
	n := 0
	for i, log := range rf.logs {
		if log.Index == rf.commitIndex {
			n = i
			break
		}
	}
	peersNum := cap(rf.peers)
	// append commit
	for n < len(rf.logs) && rf.logs[n].Index < rf.currentIndex {
		if rf.logs[n].Term != rf.currentTerm {
			n++
			continue
		}
		success := peersNum / 2
		for _, peerCommit := range rf.leaderState.matchIndex {
			if peerCommit >= rf.logs[n].Index {
				success--
			}
		}
		if success > 0 {
			break
		}
		rf.commitIndex = rf.logs[n].Index
		n++
	}
	rf.mu.Unlock()

	success := make(chan struct{})
	finish := make(chan struct{}, 1)
	go func() {
		rf.mu.Lock()
		for index, peer := range rf.peers {
			if index == int(rf.id) {
				continue
			}
			args := raftrpc.AppendEntriesArgs{
				Term:           rf.currentTerm,
				LeaderID:       rf.id,
				PrevLogTerm:    -1,
				PrevLogIndex:   -1,
				Logs:           make([]*raftrpc.Log, 0),
				LeaderCommitID: rf.commitIndex,
			}
			idx := 0
			for i, log := range rf.logs {
				if log.Index == rf.leaderState.nextIndex[int32(index)] {
					idx = i
					break
				}
			}
			if idx != 0 {
				args.PrevLogTerm = rf.logs[idx].Term
				args.PrevLogIndex = rf.logs[idx].Index
			}
			if rf.leaderState.nextIndex[int32(index)] < rf.currentIndex {

				for _, l := range rf.logs[idx:] {
					args.Logs = append(args.Logs, &raftrpc.Log{
						Term:   l.Term,
						Index:  l.Index,
						Entry:  l.Entry,
						Status: l.Status,
						Type:   int32(l.Type),
					})
				}
			}
			go rf.callAppendEntries(index, peer, &args, true, finish)
		}
		rf.mu.Unlock()
		cnt := peersNum / 2
		tot := peersNum - 1
		once := sync.Once{}
		for {
			select {
			case <-finish:
				cnt--
				tot--
				if cnt <= 0 {
					once.Do(func() {
						close(success)
					})
				}
				if tot <= 0 {
					return
				}
			}
		}
	}()

	var heartbeatOk bool
	select {
	case <-success:
		heartbeatOk = true
	case <-time.After(random.Timeout(AppendEntriesTimeout, 0)):
		heartbeatOk = false
		rf.mu.Lock()
		if rf.state != Leader {
			rf.mu.Unlock()
			return false
		}
		rf.mu.Unlock()
	}
	timeout := random.Timeout(AppendEntriesDuration, 0)
	time.Sleep(timeout)
	return heartbeatOk
}

func (rf *Raft) callAppendEntries(i int, peer raftrpc.RaftClient, args *raftrpc.AppendEntriesArgs, first bool, ch chan struct{}) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	// mem leader state
	nextIndexMem := rf.leaderState.nextIndex[int32(i)]
	matchIndexMem := rf.leaderState.matchIndex[int32(i)]
	if rf.state != Leader {
		return
	}

	rf.mu.Unlock()
	reply, err := peer.AppendEntries(context.Background(), args)
	rf.mu.Lock()
	if err != nil {
		rf.logger.Printf("can't contact instance %d through heartbeat", i)
		return
	}
	if first {
		ch <- struct{}{}
	}
	if reply.Term > rf.currentTerm {
		rf.state = Follower
		rf.stateChange <- Follower
		return
	}
	if reply.Success {
		if len(args.Logs) > 0 && rf.leaderState.matchIndex[int32(i)] < args.Logs[len(args.Logs)-1].Index {
			rf.leaderState.matchIndex[int32(i)] = args.Logs[len(args.Logs)-1].Index
			rf.leaderState.nextIndex[int32(i)] = rf.leaderState.matchIndex[int32(i)] + 1
		}
		return
	}
	// if leader state changes, it indicates that another heartbeat from next turn
	// has successfully made it to the follower, then should give this heartbeat
	if nextIndexMem != rf.leaderState.nextIndex[int32(i)] || matchIndexMem != rf.leaderState.matchIndex[int32(i)] {
		return
	}
	// index does not match follower's state, try to reduce the leader state's index
	// and send heartbeat rpc again
	logIndex := rf.leaderState.nextIndex[int32(i)] - 1
	if logIndex < 0 {
		logIndex = 0
	}
	for logIndex > 0 && rf.logs[logIndex].Term == rf.logs[logIndex-1].Term {
		logIndex--
	}

	rf.leaderState.nextIndex[int32(i)] = logIndex
	newArgs := args
	newArgs.PrevLogIndex, newArgs.PrevLogIndex = -1, -1
	newArgs.Logs = make([]*raftrpc.Log, 0)
	if rf.leaderState.nextIndex[int32(i)] > 0 {
		newArgs.PrevLogTerm = rf.logs[rf.leaderState.nextIndex[int32(i)]-1].Term
		newArgs.PrevLogIndex = rf.logs[rf.leaderState.nextIndex[int32(i)]-1].Index
	}
	if rf.leaderState.nextIndex[int32(i)] < rf.currentIndex {
		for _, l := range rf.logs[rf.leaderState.nextIndex[int32(i)]:] {
			args.Logs = append(args.Logs, &raftrpc.Log{
				Term:   l.Term,
				Index:  l.Index,
				Entry:  l.Entry,
				Status: l.Status,
			})
		}
	}
	go rf.callAppendEntries(i, peer, newArgs, false, ch)
}
