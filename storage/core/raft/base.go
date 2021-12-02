package raft

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"storage/api"
	"storage/config"
	constants "storage/constants/raft"
	"storage/helper/logger"
	"time"
)

// NewRaftInstance start a new Raft instance and return a pointer
func NewRaftInstance() *Raft {
	rpcConfig := config.GetRpcConfig()

	rf := Raft{
		id:           rpcConfig.RaftRpc.ID,
		leaderID:     rpcConfig.RaftRpc.ID,
		peers:        make([]peer, 0),
		currentTerm:  0,
		currentIndex: 0,
		commitIndex:  0,
		lastApplied:  0,
		votedFor:     constants.UnVoted,
		heartbeat:    false,
		state:        constants.Follower,
		stateChange:  make(chan constants.State),
		logs:         make([]Log, 0),
		leaderState: struct {
			nextIndex  map[int32]int32
			matchIndex map[int32]int32
		}{
			nextIndex:  make(map[int32]int32),
			matchIndex: make(map[int32]int32),
		},
	}

	for _, p := range rpcConfig.RaftPeers {
		rf.peers = append(rf.peers, peer{
			id:   p.ID,
			host: p.Host,
			port: p.Port,
		})
	}

	// start rpc server
	address := fmt.Sprintf("%s:%s", rpcConfig.RaftRpc.Host, rpcConfig.RaftRpc.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	var opts []grpc.ServerOption

	rf.rpcServer = grpc.NewServer(opts...)
	api.RegisterRaftServer(rf.rpcServer, &rf)
	err = rf.rpcServer.Serve(l)
	if err != nil {
		logger.Fatalf("Server rpc error: %v", err.Error())
	}
	logger.Printf("Serve rpc success at %s", address)

	// TODO test peer's state, only if half of the peers' rpc is served, can operations continue wo work

	// start raft
	go rf.startRaft()

	// report status periodically
	go rf.reportStatus()

	return &rf
}

// Kill to kill the raft instance
func (raft *Raft) Kill() {
	raft.state = constants.Shutdown
	raft.stateChange <- constants.Shutdown
	raft.rpcServer.GracefulStop()
}

// IsKilled check the instance is killed
func (raft *Raft) IsKilled() bool {
	return raft.state == constants.Shutdown
}

// startRaft start raft transition
func (raft *Raft) startRaft() {
	var state constants.State
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	for {
		state = <-raft.stateChange
		logger.Printf("state changed")
		cancel()
		ctx, cancel = context.WithCancel(context.Background())
		switch state {
		case constants.Follower:
			go raft.follower(ctx)
		case constants.Candidate:
			go raft.candidate(ctx)
		case constants.Leader:
			go raft.leader(ctx)
		case constants.Shutdown:
			cancel()
			return
		}
	}
}

// reportStatus add log to see instance's status
func (raft *Raft) reportStatus() {
	for {
		raft.mu.Lock()
		if raft.state == constants.Shutdown {
			raft.mu.Unlock()
			return
		}
		logger.Printf("{ term: %d, index: %d }", raft.currentTerm, raft.currentIndex)
		raft.mu.Unlock()
		time.Sleep(constants.StatusLoggerTimeout)
	}
}
