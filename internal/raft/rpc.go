package raft

import (
	"configStorage/api/raftrpc"
	"context"
	"time"
)

func (rf *Raft) RequestVote(ctx context.Context, args *raftrpc.RequestVoteArgs) (reply *raftrpc.RequestVoteReply, err error) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	reply = &raftrpc.RequestVoteReply{}
	//raft.logger.Printf("receive vote request from instance %d [term: %d, index: %d]", args.CandidateID, args.Term, args.LastLogIndex)
	reply.Term = rf.currentTerm
	// term out of time, return false
	if args.Term < rf.currentTerm {
		rf.logger.Printf("candidate's term is lower then mine, deny it")
		reply.VoteGranted = false
		return reply, nil
	}

	if len(rf.logs) > 0 {
		lastLogTerm := rf.logs[len(rf.logs)-1].Term
		if lastLogTerm > args.LastLogTerm ||
			(lastLogTerm == args.LastLogTerm && rf.logs[len(rf.logs)-1].Index > args.LastLogIndex) {
			rf.logger.Printf("candidate's logs are older then mine, deny it")
			reply.VoteGranted = false
			return reply, nil
		}
	}

	if rf.votedFor == UnVoted || rf.votedFor == args.CandidateID {
		rf.logger.Printf("voted for instance %d", args.CandidateID)
		rf.votedFor = args.CandidateID
		reply.VoteGranted = true
		// persist
		rf.persist()
		return reply, nil
	}
	reply.VoteGranted = false
	return reply, nil
}

func (rf *Raft) AppendEntries(ctx context.Context, args *raftrpc.AppendEntriesArgs) (reply *raftrpc.AppendEntriesReply, err error) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply = &raftrpc.AppendEntriesReply{}
	//raft.logger.Printf("receive heartbeat from leader %d [ term : %d, index: %d ]", args.LeaderID, args.Term, args.PrevLogIndex)
	reply.Term = rf.currentTerm

	if rf.currentTerm > args.Term {
		rf.logger.Printf("current term is bigger then leader's term, deny heartbeat")
		reply.Success = false
		return reply, nil
	} else if rf.state == Candidate {
		// if instance is a candidate, deny the request
		// and inform the leader's lost by add term id
		if rf.currentTerm == args.Term {
			rf.logger.Printf("deny leader's heartbeat for start a new round of election")
			reply.Term = rf.currentTerm + 1
			reply.Success = false
			return reply, nil
		}
		// get a new leader, give up this election
		rf.logger.Printf("give up my election, follow leader's state")
		rf.state = Follower
		rf.stateChange <- Follower
	}

	// set heartbeat status
	rf.currentTerm = args.Term
	rf.votedFor = UnVoted
	rf.heartbeat = true
	rf.leaderID = args.LeaderID

	if args.PrevLogIndex != NonLogIndex {
		// can't match leader's state last entry index
		if rf.currentIndex <= args.PrevLogIndex ||
			rf.logs[args.PrevLogIndex].Term != args.PrevLogTerm {
			reply.Success = false
			return reply, nil
		}
	}

	// finish testify, follow leader's state
	for i, log := range args.Logs {
		if log.Index <= rf.commitIndex {
			continue
		}
		if rf.currentIndex > log.Index {
			// term not match, follow leader's logs
			if rf.logs[i].Term != log.Term {
				rf.logs = rf.logs[:i]
				for _, l := range args.Logs[i:] {
					rf.logs = append(rf.logs, Log{
						Entry:  l.Entry,
						Term:   l.Term,
						Index:  l.Index,
						Status: l.Status,
						Type:   LogType(l.Type),
					})
				}
				break
			} else {
				continue
			}
		}
		rf.logs = append(rf.logs, Log{
			Entry:  log.Entry,
			Term:   log.Term,
			Index:  log.Index,
			Status: log.Status,
			Type:   LogType(log.Type),
		})
		rf.currentIndex++
	}

	if len(rf.logs) > 0 {
		rf.currentIndex = rf.logs[len(rf.logs)-1].Index + 1
	}

	// follow leader's commit id
	if args.LeaderCommitID > rf.commitIndex {
		lastIndex := rf.currentIndex - 1
		if lastIndex < args.LeaderCommitID {
			rf.commitIndex = lastIndex
		} else {
			rf.commitIndex = args.LeaderCommitID
		}
	}

	// persist
	rf.persist()
	reply.Success = true
	return reply, nil
}

// NewEntry append a new entry to leader's log
func (rf *Raft) NewEntry(ctx context.Context, args *raftrpc.NewEntryArgs) (reply *raftrpc.NewEntryReply, err error) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	now := time.Now()
	// if not a leader, return leader's msg
	reply = &raftrpc.NewEntryReply{}
	reply.LeaderID = rf.leaderID
	if rf.state != Leader {
		reply.Success = false
		reply.Msg = "instance not a leader"
		return reply, nil
	}

	logIndex := rf.currentIndex
	rf.currentIndex = rf.currentIndex + 1
	// new log
	log := Log{
		Entry:  args.Entry,
		Term:   rf.currentTerm,
		Type:   LogType(args.Type),
		Index:  logIndex,
		Status: true,
	}
	rf.logs = append(rf.logs, log)
	rf.mu.Unlock()

	rf.logger.Printf("new log entry at term %d index %d", log.Term, log.Index)
	// check 40 times, for total 2s
	for i := 0; i < 40; i++ {
		time.Sleep(NewEntryTimeout)
		rf.mu.Lock()
		// if commit index > log index and commit success
		if rf.commitIndex >= logIndex && rf.logs[logIndex].Status == true {
			// calculate commit msg duration
			d := time.Now().Sub(now)
			rf.logger.Printf("log %d committed, using %d micro seconds", log.Index, d.Microseconds())
			if len(rf.msgCommitTime) == 100 {
				rf.msgCommitTime = rf.msgCommitTime[1:]
			}
			rf.msgCommitTime = append(rf.msgCommitTime, d)
			reply.Success = true
			reply.Msg = "ok"
			return reply, nil
		}
		// commit fail
		if rf.logs[logIndex].Status == false {
			rf.logger.Printf("log %d executed error", log.Index)
			reply.Success = false
			reply.Msg = "entry commit error"
			return reply, nil
		}
		rf.mu.Unlock()
	}
	rf.logger.Printf("log %d process timeout", log.Index)
	reply.Success = false
	reply.Msg = "Timeout"
	return reply, nil
}

func (rf *Raft) GetValue(ctx context.Context, args *raftrpc.GetValueArgs) (reply *raftrpc.GetValueReply, err error) {
	reply = &raftrpc.GetValueReply{}
	v, err := rf.storage.Get(args.GetKey())
	reply.Value = v
	return reply, nil
}

func (rf *Raft) GetPrefixConfigs(ctx context.Context, args *raftrpc.GetPrefixConfigArgs) (reply *raftrpc.GetPrefixConfigReply, err error) {
	reply = &raftrpc.GetPrefixConfigReply{}
	v := rf.storage.PrefixAll(args.Prefix)
	reply.Config = *v
	return reply, nil
}
