package raft

import (
	"configStorage/api/raftrpc"
	"configStorage/internal/config"
	"configStorage/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strings"
	"time"
)

// NewRaftInstance start a new Raft instance and return a pointer
func NewRaftInstance(rpcConfig config.Raft) *Raft {
	rf := Raft{
		id:           rpcConfig.RaftRpc.ID,
		leaderID:     rpcConfig.RaftRpc.ID,
		peers:        make([]raftrpc.RaftClient, len(rpcConfig.RaftPeers)),
		currentTerm:  0,
		currentIndex: 0,
		commitIndex:  -1,
		lastApplied:  -1,
		votedFor:     UnVoted,
		heartbeat:    false,
		state:        Follower,
		stateChange:  make(chan State, 1),
		logs:         make([]Log, 0),
		logger:       logger.NewLogger([]interface{}{rpcConfig.RaftRpc.ID}, rpcConfig.LogPrefix),
		cfg:          rpcConfig,
		storage:      NewRaftStorage(),
		leaderState: struct {
			nextIndex  map[int32]int32
			matchIndex map[int32]int32
		}{
			nextIndex:  make(map[int32]int32),
			matchIndex: make(map[int32]int32),
		},
	}
	return &rf
}

func (rf *Raft) Start() {

	// start rpc server
	address := fmt.Sprintf("%s:%s", rf.cfg.RaftRpc.Host, rf.cfg.RaftRpc.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}

	c_address := fmt.Sprintf("%s:%s", rf.cfg.RaftRpc.Host, rf.cfg.RaftRpc.CPort)
	cl, err := net.Listen("tcp", c_address)
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	var sOpts []grpc.ServerOption

	rf.rpcServer = grpc.NewServer(sOpts...)
	raftrpc.RegisterRaftServer(rf.rpcServer, rf)

	var cOpts []grpc.ServerOption
	rf.stateServer = grpc.NewServer(cOpts...)
	raftrpc.RegisterStateServer(rf.stateServer, rf)

	// state server start
	go func() {
		err1 := rf.stateServer.Serve(cl)
		if err1 != nil {
			rf.logger.Fatalf("Start rpc server error: %v", err1.Error())
		}
	}()

	go func() {
		rf.logger.Printf("waiting for peers' rpc server to start up")
		time.Sleep(StartUpTimeout)

		cOpts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		failNum := 0
		for _, p := range rf.cfg.RaftPeers {
			if p.ID != rf.id {
				serverAddr := fmt.Sprintf("%s:%s", p.Host, p.Port)
				conn, err := grpc.Dial(serverAddr, cOpts...)
				if err != nil {
					rf.logger.Printf("open connection with id %d error, addr: %s, error: %v", p.ID, serverAddr, err.Error())
					failNum++
					continue
				}
				client := raftrpc.NewRaftClient(conn)
				rf.peers[p.ID] = client
			}
		}
		if failNum > len(rf.cfg.RaftPeers)/2 {
			//rf.Kill()
			rf.logger.Fatalf("over half of the peer client is closed")
		}

		// start raft
		go rf.startRaft()

		go rf.checkCommit()
		// set state
		rf.stateChange <- Follower

		// report status periodically
		go rf.reportStatus()
	}()

	rf.logger.Printf("Serving rpc at %s", address)
	err = rf.rpcServer.Serve(l)
	if err != nil {
		rf.logger.Fatalf("Server rpc error: %v", err.Error())
	}
}

// Kill to kill the raft instance
func (rf *Raft) Kill() {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.state = Shutdown
	rf.stateChange <- Shutdown
	rf.rpcServer.GracefulStop()
	rf.stateServer.GracefulStop()
}

func (rf *Raft) Restart(host string, port string) {
	address := fmt.Sprintf("%s:%s", host, port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	var sOpts []grpc.ServerOption

	rf.rpcServer = grpc.NewServer(sOpts...)
	raftrpc.RegisterRaftServer(rf.rpcServer, rf)
	go func() {
		err = rf.rpcServer.Serve(l)
	}()
	if err != nil {
		rf.logger.Fatalf("Server rpc error: %v", err.Error())
	}
	rf.logger.Printf("restart Serve rpc success at %s", address)
	// start raft
	go rf.startRaft()

	// set state
	rf.stateChange <- Follower
	rf.state = Follower

	// report status periodically
	go rf.reportStatus()

}

// IsKilled check the instance is killed
func (rf *Raft) IsKilled() bool {
	return rf.state == Shutdown
}

// startRaft start raft transition
func (rf *Raft) startRaft() {
	var state State
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	for {
		state = <-rf.stateChange
		rf.logger.Printf("state changed")
		cancel()
		ctx, cancel = context.WithCancel(context.Background())
		switch state {
		case Follower:
			go rf.follower(ctx)
		case Candidate:
			go rf.candidate(ctx)
		case Leader:
			go rf.leader(ctx)
		case Shutdown:
			cancel()
			return
		}
	}
}

// reportStatus add log to see instance's status
func (rf *Raft) reportStatus() {
	for {
		rf.mu.Lock()
		if rf.state == Shutdown {
			rf.mu.Unlock()
			return
		}
		rf.logger.Printf("{ state: %v, term: %d, index: %d }", rf.state, rf.currentTerm, rf.currentIndex)
		rf.mu.Unlock()
		time.Sleep(StatusLoggerTimeout)
	}
}

func (rf *Raft) checkCommit() {
	for {
		rf.mu.Lock()
		if rf.state == Shutdown {
			rf.mu.Unlock()
			return
		}
		if rf.lastApplied < rf.commitIndex {
			for i, log := range rf.logs {
				if log.Index > rf.lastApplied && log.Index <= rf.commitIndex {
					command := string(log.Entry)
					words := strings.Split(command, " ")
					if words[0] == "set" {
						if len(words) <= 2 {
							rf.logger.Printf("entry format error")
							rf.logs[i].Status = false
						}
						key := words[1]
						value := strings.Join(words[2:], " ")
						rf.storage.Set(key, value)
					} else if words[0] == "del" {
						if len(words) != 2 {
							rf.logger.Printf("entry format error")
							rf.logs[i].Status = false
						}
						key := words[1]
						if rf.storage.Del(key) != nil {
							rf.logger.Printf("error del value, %v", command)
							rf.logs[i].Status = false
						}
					} else {
						rf.logger.Printf("entry format error")
						rf.logs[i].Status = false
					}
					rf.lastApplied = log.Index
				}
			}
		}
		rf.mu.Unlock()
		time.Sleep(NewEntryTimeout)
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
