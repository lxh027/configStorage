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
func NewRaftInstance(rpcConfig config.RpcConfig) *Raft {
	rf := Raft{
		id:           rpcConfig.RaftRpc.ID,
		leaderID:     rpcConfig.RaftRpc.ID,
		peers:        make([]api.RaftClient, len(rpcConfig.RaftPeers)),
		currentTerm:  0,
		currentIndex: 0,
		commitIndex:  0,
		lastApplied:  0,
		votedFor:     constants.UnVoted,
		heartbeat:    false,
		state:        constants.Follower,
		stateChange:  make(chan constants.State, 1),
		logs:         make([]Log, 0),
		logger:       logger.NewLogger(rpcConfig.RaftRpc.ID),
		cfg:          rpcConfig,
		leaderState: struct {
			nextIndex  map[int32]int32
			matchIndex map[int32]int32
		}{
			nextIndex:  make(map[int32]int32),
			matchIndex: make(map[int32]int32),
		},
	}

	// start rpc server
	address := fmt.Sprintf("%s:%s", rpcConfig.RaftRpc.Host, rpcConfig.RaftRpc.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	var sOpts []grpc.ServerOption

	rf.rpcServer = grpc.NewServer(sOpts...)
	api.RegisterRaftServer(rf.rpcServer, &rf)

	go func() {
		err = rf.rpcServer.Serve(l)
	}()
	if err != nil {
		rf.logger.Fatalf("Server rpc error: %v", err.Error())
	}
	rf.logger.Printf("Serve rpc success at %s", address)

	rf.logger.Printf("waiting for peers' rpc server to start up")
	time.Sleep(constants.StartUpTimeout)

	cOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	failNum := 0
	for _, p := range rpcConfig.RaftPeers {
		if p.ID != rf.id {
			serverAddr := fmt.Sprintf("%s:%s", p.Host, p.Port)
			conn, err := grpc.Dial(serverAddr, cOpts...)
			if err != nil {
				rf.logger.Printf("open connection with id %d error, addr: %s, error: %v", p.ID, serverAddr, err.Error())
				failNum++
				continue
			}
			client := api.NewRaftClient(conn)
			rf.peers[p.ID] = client
		}
	}
	if failNum > len(rpcConfig.RaftPeers)/2 {
		//rf.Kill()
		rf.logger.Fatalf("over half of the peer client is closed")
	}

	// start raft
	go rf.startRaft()

	// set state
	rf.stateChange <- constants.Follower

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

func (raft *Raft) Restart(host string, port string) {
	address := fmt.Sprintf("%s:%s", host, port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		raft.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	var sOpts []grpc.ServerOption

	raft.rpcServer = grpc.NewServer(sOpts...)
	api.RegisterRaftServer(raft.rpcServer, raft)
	go func() {
		err = raft.rpcServer.Serve(l)
	}()
	if err != nil {
		raft.logger.Fatalf("Server rpc error: %v", err.Error())
	}
	raft.logger.Printf("restart Serve rpc success at %s", address)
	// start raft
	go raft.startRaft()

	// set state
	raft.stateChange <- constants.Follower
	raft.state = constants.Follower

	// report status periodically
	go raft.reportStatus()

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
		raft.logger.Printf("state changed")
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
		raft.logger.Printf("{ state: %v, term: %d, index: %d }", raft.state, raft.currentTerm, raft.currentIndex)
		raft.mu.Unlock()
		time.Sleep(constants.StatusLoggerTimeout)
	}
}

/*func (raft *Raft) getConn(id int) *grpc.ClientConn {
	raft.mu.Lock()
	defer raft.mu.Unlock()
	if raft.peers[id] == nil || raft.peers[id].GetState() != connectivity.Ready{
		cOpts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		serverAddr := fmt.Sprintf("%s:%s", raft.cfg.RaftPeers[id].Host, raft.cfg.RaftPeers[id].Port)

		conn, err := grpc.Dial(serverAddr, cOpts...)
		if err != nil {
			raft.logger.Printf("reconnection with id %d error, addr: %s, error: %v", id, serverAddr, err.Error())
		}
		_ = raft.peers[id].Close()
		raft.peers[id] = conn
	}
	return raft.peers[id]
}*/
