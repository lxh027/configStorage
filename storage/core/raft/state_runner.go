package raft

import (
	"context"
	constants "storage/constants/raft"
	"storage/helper/logger"
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
	for {
		select {
		case <-ctx.Done():
			logger.Printf("finish state as leader")
			return
		case <-time.After(constants.LeaderRentTimeout):
			// if state is a leader, then set a rent for seconds
			// when rent is due, finish the instance as a leader,
			// then start a new round of election
			logger.Printf("rent is due as a leader")
			return
		default:
			raft.loopLeader()
		}
	}
}

func (raft *Raft) loopFollower() {
	// TODO a loop for follower
}

func (raft *Raft) loopCandidate() {
	// TODO a loop for candidate
}

func (raft *Raft) loopLeader() {
	// TODO a loop for leader
}
