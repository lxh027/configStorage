package raft

import (
	"context"
	"storage/api"
	constants "storage/constants/raft"
	"time"
)

func (raft *Raft) RequestVote(ctx context.Context, args *api.RequestVoteArgs) (reply *api.RequestVoteReply, err error) {
	raft.mu.Lock()
	defer raft.mu.Unlock()
	reply = &api.RequestVoteReply{}
	//raft.logger.Printf("receive vote request from instance %d [term: %d, index: %d]", args.CandidateID, args.Term, args.LastLogIndex)
	reply.Term = raft.currentTerm
	// term out of time, return false
	if args.Term < raft.currentTerm {
		raft.logger.Printf("candidate's term is lower then mine, deny it")
		reply.VoteGranted = false
		return reply, nil
	}

	if raft.currentIndex > 0 {
		lastLogTerm := raft.logs[raft.currentIndex-1].Term
		if lastLogTerm > args.LastLogTerm ||
			(lastLogTerm == args.LastLogTerm && raft.logs[raft.currentIndex-1].Index > args.LastLogIndex) {
			raft.logger.Printf("candidate's logs are older then mine, deny it")
			reply.VoteGranted = false
			return reply, nil
		}
	}

	if raft.votedFor == constants.UnVoted || raft.votedFor == args.CandidateID {
		raft.logger.Printf("voted for instance %d", args.CandidateID)
		raft.votedFor = args.CandidateID
		reply.VoteGranted = true
		// TODO: persist

		return reply, nil
	}
	reply.VoteGranted = false
	return reply, nil
}

func (raft *Raft) AppendEntries(ctx context.Context, args *api.AppendEntriesArgs) (reply *api.AppendEntriesReply, err error) {
	raft.mu.Lock()
	defer raft.mu.Unlock()

	reply = &api.AppendEntriesReply{}
	//raft.logger.Printf("receive heartbeat from leader %d [ term : %d, index: %d ]", args.LeaderID, args.Term, args.PrevLogIndex)
	reply.Term = raft.currentTerm

	if raft.currentTerm > args.Term {
		raft.logger.Printf("current term is bigger then leader's term, deny heartbeat")
		reply.Success = false
		return reply, nil
	} else if raft.state == constants.Candidate {
		// if instance is a candidate, deny the request
		// and inform the leader's lost by add term id
		if raft.currentTerm == args.Term {
			raft.logger.Printf("deny leader's heartbeat for start a new round of election")
			reply.Term = raft.currentIndex + 1
			reply.Success = false
			return reply, nil
		}
		// get a new leader, give up this election
		raft.logger.Printf("give up my election, follow leader's state")
		raft.state = constants.Follower
		raft.stateChange <- constants.Follower
	}

	// set heartbeat status
	raft.currentTerm = args.Term
	raft.votedFor = constants.UnVoted
	raft.heartbeat = true

	if args.PrevLogIndex != constants.NonLogIndex {
		// can't match leader's state last entry index
		if raft.currentIndex <= args.PrevLogIndex ||
			raft.logs[args.PrevLogIndex].Term != args.PrevLogTerm {
			reply.Success = false
			return reply, nil
		}
	}

	// finish testify, follow leader's state
	for i, log := range args.Logs {
		if log.Index <= raft.commitIndex {
			continue
		}
		if raft.currentIndex > log.Index {
			// term not match, follow leader's logs
			if raft.logs[log.Index].Term != log.Term {
				raft.logs = raft.logs[:log.Index]
				for _, l := range args.Logs[i:] {
					raft.logs = append(raft.logs, Log{
						Entry:  l.Entry,
						Term:   l.Term,
						Index:  l.Index,
						Status: l.Status,
					})
				}
				break
			} else {
				continue
			}
		}
		raft.logs = append(raft.logs, Log{
			Entry:  log.Entry,
			Term:   log.Term,
			Index:  log.Index,
			Status: log.Status,
		})
		raft.currentIndex++
	}

	raft.currentIndex = int32(len(raft.logs))

	// follow leader's commit id
	if args.LeaderCommitID > raft.commitIndex {
		lastIndex := raft.currentIndex - 1
		if lastIndex < args.LeaderCommitID {
			raft.commitIndex = lastIndex
		} else {
			raft.commitIndex = args.LeaderCommitID
		}
	}

	// TODO persist

	reply.Success = true
	return reply, nil
}

// NewEntry append a new entry to leader's log
func (raft *Raft) NewEntry(ctx context.Context, args *api.NewEntryArgs) (reply *api.NewEntryReply, err error) {
	raft.mu.Lock()
	defer raft.mu.Unlock()

	// if not a leader, return leader's msg
	reply = &api.NewEntryReply{}
	reply.LeaderID = raft.leaderID
	if raft.state != constants.Leader {
		reply.Success = false
		reply.Msg = "instance not a leader"
		return reply, nil
	}

	logIndex := raft.currentIndex
	raft.currentIndex = raft.currentIndex + 1
	// new log
	log := Log{
		Entry:  args.Entry,
		Term:   raft.currentTerm,
		Index:  logIndex,
		Status: true,
	}
	raft.logs = append(raft.logs, log)
	raft.mu.Unlock()

	raft.logger.Printf("new log entry at term %d index %d", log.Term, log.Index)
	// check 4 times, for total 2s
	for i := 0; i < 4; i++ {
		time.Sleep(constants.NewEntryTimeout)
		raft.mu.Lock()
		// if commit index > log index and commit success
		if raft.commitIndex >= logIndex && raft.logs[logIndex].Status == true {
			raft.logger.Printf("log %d committed", log.Index)
			reply.Success = true
			reply.Msg = "ok"
			return reply, nil
		}
		// commit fail
		if raft.logs[logIndex].Status == false {
			raft.logger.Printf("log %d executed error", log.Index)
			reply.Success = false
			reply.Msg = "entry commit error"
			return reply, nil
		}
		raft.mu.Unlock()
	}
	raft.logger.Printf("log %d process timeout", log.Index)
	reply.Success = false
	reply.Msg = "Timeout"
	return reply, nil
}
