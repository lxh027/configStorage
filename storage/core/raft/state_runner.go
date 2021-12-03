package raft

import (
	"context"
	"storage/api"
	constants "storage/constants/raft"
	"storage/helper"
	"storage/helper/logger"
	"sync"
	"time"
)

func (raft *Raft) follower(ctx context.Context) {
	logger.Printf("start state as follower")
	for {
		select {
		case <-ctx.Done():
			logger.Printf("finish state as follower")
			return
		default:
			raft.loopFollower()
		}
	}
}

func (raft *Raft) candidate(ctx context.Context) {
	logger.Printf("start state ad candidate")
	for {
		select {
		case <-ctx.Done():
			logger.Printf("finish state as candidate")
			return
		default:
			raft.loopCandidate()
		}
	}
}

func (raft *Raft) leader(ctx context.Context) {
	logger.Printf("start state ad leader")
	var heartbeatOk = false
	rentDue := make(chan struct{})

	go func() {
		// if state is a leader, then set a rent for seconds
		// when rent is due, finish the instance as a leader,
		// then start a new round of election
		for {
			select {
			case <-time.After(constants.LeaderRentTimeout):
				if !heartbeatOk {
					close(rentDue)
					return
				}
				heartbeatOk = false
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			logger.Printf("finish state as leader")
			return
		case <-rentDue:
			logger.Printf("rent is due as a leader")
			raft.mu.Lock()
			raft.state = constants.Follower
			raft.stateChange <- constants.Follower
			raft.mu.Unlock()
			return
		default:
			heartbeatOk = raft.loopLeader()
		}
	}
}

func (raft *Raft) loopFollower() {
	defer raft.mu.Unlock()
	raft.mu.Lock()
	// reset info as a follower
	raft.votedFor = constants.UnVoted
	raft.heartbeat = false
	raft.mu.Unlock()

	timeout := helper.RandomTimeout(constants.HeartBeatTimeout, constants.RandTimeout)
	time.Sleep(timeout)

	raft.mu.Lock()
	// if after sleep and not receive from leader
	// and haven't voted for anyone, then start a new election
	if !raft.heartbeat && raft.state == constants.Follower &&
		(raft.votedFor == constants.UnVoted || raft.votedFor == raft.id) {
		raft.state = constants.Candidate
		raft.stateChange <- constants.Candidate
	}
	raft.mu.Unlock()
}

func (raft *Raft) loopCandidate() {
	defer raft.mu.Unlock()
	raft.mu.Lock()

	if raft.state != constants.Candidate ||
		(raft.votedFor != raft.id && raft.votedFor != constants.UnVoted) {
		return
	}

	peersNum := cap(raft.peers)
	wg := sync.WaitGroup{}
	wg.Add((peersNum + 1) / 2)
	success := make(chan struct{})
	raft.mu.Unlock()

	go func() {
		raft.mu.Lock()
		defer raft.mu.Unlock()
		for i, peer := range raft.peers {
			if i != int(raft.id) {
				args := api.RequestVoteArgs{
					Term:         raft.currentTerm,
					CandidateID:  raft.id,
					LastLogIndex: 0,
					LastLogTerm:  0,
				}
				if raft.currentIndex > 0 {
					args.LastLogIndex = raft.logs[raft.currentIndex-1].Index
					args.LastLogTerm = raft.logs[raft.currentIndex-1].Term
				}
				// send rpc request for votes
				go func(peerIndex int, peer api.RaftClient, args *api.RequestVoteArgs) {
					defer raft.mu.Unlock()
					reply, err := peer.RequestVote(context.Background(), args)
					if err != nil {
						logger.Printf("fail to send rpc request to id %d: %v", peerIndex, err.Error())
						return
					}
					raft.mu.Lock()
					if reply.Term > raft.currentTerm {
						raft.currentTerm = reply.Term
						raft.state = constants.Follower
						raft.votedFor = constants.UnVoted
						raft.stateChange <- constants.Follower
						return
					}
					raft.mu.Unlock()
					if reply.VoteGranted {
						logger.Printf("receive vote from peer %d", peerIndex)
						wg.Done()
						return
					}
				}(i, peer, &args)
			}
		}
		raft.mu.Unlock()
		wg.Wait()
		close(success)
	}()

	select {
	case <-success:
		raft.mu.Lock()
		if raft.state != constants.Candidate {
			return
		}
		raft.currentTerm++
		for index := 0; index < peersNum; index++ {
			raft.leaderState.matchIndex[int32(index)] = raft.currentIndex
			raft.leaderState.nextIndex[int32(index)] = 0
		}
		raft.state = constants.Leader
		raft.stateChange <- constants.Leader
		raft.mu.Unlock()
	case <-time.After(helper.RandomTimeout(constants.ElectionTimeout, 0)):
		raft.mu.Lock()
		if raft.state != constants.Candidate {
			return
		}
		raft.votedFor = constants.UnVoted
		raft.mu.Unlock()
		// sleep and wait for another election
		timeout := helper.RandomTimeout(constants.LatencyTimeout, constants.RandTimeout)
		time.Sleep(timeout)
	}
}

// loopLeader main func when state is leader
// return true if successfully contact all the other peers
func (raft *Raft) loopLeader() bool {
	defer raft.mu.Unlock()
	raft.mu.Lock()
	n := raft.commitIndex + 1
	peersNum := cap(raft.peers)
	// append commit
	for n >= 0 && n < raft.currentIndex {
		if raft.logs[n].Term != raft.currentTerm {
			continue
		}
		success := peersNum / 2
		for _, peerCommit := range raft.leaderState.matchIndex {
			if peerCommit >= n {
				success--
			}
		}
		if success > 0 {
			break
		}
		raft.commitIndex = n
		n++
	}
	raft.mu.Unlock()

	// TODO persist
	wg := sync.WaitGroup{}
	wg.Add((peersNum + 1) / 2)
	var success chan struct{}
	go func() {
		raft.mu.Lock()
		defer raft.mu.Unlock()
		for index, peer := range raft.peers {
			if index == int(raft.id) {
				continue
			}
			args := api.AppendEntriesArgs{
				Term:           raft.currentTerm,
				LeaderID:       raft.id,
				PrevLogTerm:    -1,
				PrevLogIndex:   -1,
				Logs:           make([]*api.Log, 0),
				LeaderCommitID: raft.commitIndex,
			}
			if raft.leaderState.nextIndex[int32(index)] != 0 {
				args.PrevLogTerm = raft.logs[raft.leaderState.nextIndex[int32(index)]-1].Term
				args.PrevLogIndex = raft.logs[raft.leaderState.nextIndex[int32(index)]-1].Index
			}
			if raft.leaderState.nextIndex[int32(index)] < raft.currentIndex {
				for _, l := range raft.logs[raft.leaderState.nextIndex[int32(index)]:] {
					args.Logs = append(args.Logs, &api.Log{
						Term:   l.Term,
						Index:  l.Index,
						Entry:  l.Entry,
						Status: l.Status,
					})
				}
			}
			go raft.callAppendEntries(index, peer, &args, true, &wg)
		}
		raft.mu.Unlock()
		wg.Wait()
		close(success)
	}()

	var heartbeatOk bool
	select {
	case <-success:
		heartbeatOk = true
	case <-time.After(helper.RandomTimeout(constants.AppendEntriesTimeout, 0)):
		heartbeatOk = false
		raft.mu.Lock()
		if raft.state != constants.Leader {
			return false
		}
		raft.mu.Unlock()
		timeout := helper.RandomTimeout(constants.AppendEntriesDuration, 0)
		time.Sleep(timeout)
	}
	return heartbeatOk
}

func (raft *Raft) callAppendEntries(i int, client api.RaftClient, args *api.AppendEntriesArgs, first bool, wg *sync.WaitGroup) {
	raft.mu.Lock()
	defer raft.mu.Unlock()
	// mem leader state
	nextIndexMem := raft.leaderState.nextIndex[int32(i)]
	matchIndexMem := raft.leaderState.matchIndex[int32(i)]
	if raft.state != constants.Leader {
		return
	}

	raft.mu.Unlock()
	reply, err := client.AppendEntries(context.Background(), args)
	if err != nil {
		logger.Printf("can't contact instance %d through heartbeat", i)
		return
	}
	if first {
		wg.Done()
	}
	raft.mu.Lock()
	if reply.Term > raft.currentTerm {
		raft.state = constants.Follower
		raft.stateChange <- constants.Follower
		return
	}
	if reply.Success {
		if len(args.Logs) > 0 && raft.leaderState.matchIndex[int32(i)] < args.Logs[len(args.Logs)-1].Index {
			raft.leaderState.matchIndex[int32(i)] = args.Logs[len(args.Logs)-1].Index
			raft.leaderState.nextIndex[int32(i)] = raft.leaderState.matchIndex[int32(i)] + 1
		}
		return
	}
	// if leader state changes, it indicates that another heartbeat from next turn
	// has successfully made it to the follower, then should give this heartbeat
	if nextIndexMem != raft.leaderState.nextIndex[int32(i)] || matchIndexMem != raft.leaderState.matchIndex[int32(i)] {
		return
	}
	// index does not match follower's state, try to reduce the leader state's index
	// and send heartbeat rpc again
	logIndex := raft.leaderState.nextIndex[int32(i)] - 1
	if logIndex < 0 {
		logIndex = 0
	}
	for logIndex > 0 && raft.logs[logIndex].Term == raft.logs[logIndex-1].Term {
		logIndex--
	}

	raft.leaderState.nextIndex[int32(i)] = logIndex
	newArgs := args
	newArgs.PrevLogIndex, newArgs.PrevLogIndex = -1, -1
	newArgs.Logs = make([]*api.Log, 0)
	if raft.leaderState.nextIndex[int32(i)] > 0 {
		newArgs.PrevLogTerm = raft.logs[raft.leaderState.nextIndex[int32(i)]-1].Term
		newArgs.PrevLogIndex = raft.logs[raft.leaderState.nextIndex[int32(i)]-1].Index
	}
	if raft.leaderState.nextIndex[int32(i)] < raft.currentIndex {
		for _, l := range raft.logs[raft.leaderState.nextIndex[int32(i)]:] {
			args.Logs = append(args.Logs, &api.Log{
				Term:   l.Term,
				Index:  l.Index,
				Entry:  l.Entry,
				Status: l.Status,
			})
		}
	}
	raft.mu.Unlock()
	go raft.callAppendEntries(i, client, newArgs, false, wg)
}
