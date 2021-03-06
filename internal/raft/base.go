package raft

import (
	"bytes"
	"configStorage/api/raftrpc"
	"configStorage/pkg/config"
	"configStorage/pkg/logger"
	"configStorage/pkg/redis"
	"context"
	"encoding/json"
	"fmt"
	"github.com/HuyuYasumi/kvuR/labgob"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/shirou/gopsutil/mem"
	"google.golang.org/grpc"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

// NewRaftInstance start a new Raft instance and return a pointer
func NewRaftInstance(rpcConfig Config, rfConfig RfConfig, redisConfig config.Redis) *Raft {
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
		persister:    NewPersister(),
		redisClient:  nil,
		redisCfg:     redisConfig,
		cfg:          rpcConfig,
		raftCfg:      rfConfig,
		storage:      NewNamespaceStorage(),
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

func (rf *Raft) Start(md5 string) {
	rf.cfgVersion = md5
	// start rpc server
	address := fmt.Sprintf("%s:%s", rf.raftCfg.Host, rf.raftCfg.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}

	c_address := fmt.Sprintf("%s:%s", rf.raftCfg.Host, rf.raftCfg.CPort)
	cl, err := net.Listen("tcp", c_address)
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	sOpts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger.NewZapLogger()))),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger.NewZapLogger()))),
	}

	rf.rpcServer = grpc.NewServer(sOpts...)
	raftrpc.RegisterRaftServer(rf.rpcServer, rf)

	rf.stateServer = grpc.NewServer(sOpts...)
	raftrpc.RegisterStateServer(rf.stateServer, rf)

	// connect redis
	go func() {
		for {
			var redisErr error
			rf.redisClient, redisErr = redis.NewRedisClient(&rf.redisCfg)
			if redisErr == nil {
				break
			}
			rf.logger.Printf("error connecting redis: %v", redisErr.Error())
		}
	}()

	// state server start
	go func() {
		rf.logger.Printf("Serving client rpc at %s", c_address)
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

		//rf.readPersist()

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
	rf.persist()
	rf.snapshot()
	rf.state = Shutdown
	rf.stateChange <- Shutdown
	go rf.rpcServer.Stop()
	//go rf.stateServer.Stop()
}

func (rf *Raft) Restart() {
	// start rpc server
	address := fmt.Sprintf("%s:%s", rf.raftCfg.Host, rf.raftCfg.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	if err != nil {
		rf.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	sOpts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger.NewZapLogger()))),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger.NewZapLogger()))),
	}

	rf.rpcServer = grpc.NewServer(sOpts...)
	raftrpc.RegisterRaftServer(rf.rpcServer, rf)

	rf.stateServer = grpc.NewServer(sOpts...)
	raftrpc.RegisterStateServer(rf.stateServer, rf)

	// connect redis
	go func() {
		for {
			var redisErr error
			rf.redisClient, redisErr = redis.NewRedisClient(&rf.redisCfg)
			if redisErr == nil {
				break
			}
			rf.logger.Printf("error connecting redis: %v", redisErr.Error())
		}
	}()

	go func() {
		rf.logger.Printf("Serving raft rpc at %s", address)
		err = rf.rpcServer.Serve(l)
		if err != nil {
			rf.logger.Fatalf("Start rpc server error: %v", err.Error())
		}
	}()
	rf.readPersist()
	// start raft
	go rf.checkCommit()
	go rf.startRaft()

	// set state
	rf.stateChange <- Follower
	rf.state = Follower

	// report status periodically
	// go rf.reportStatus()

}

// IsKilled check the instance is killed
func (rf *Raft) IsKilled() bool {
	return rf.state == Shutdown
}

func (rf *Raft) Status() State {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return rf.state
}

// startRaft start raft transition
func (rf *Raft) startRaft() {
	var state State
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	var now = time.Now()
	var lastState = Follower
	for {
		state = <-rf.stateChange
		rf.persist()
		rf.snapshot()
		rf.logger.Printf("state changed")
		cancel()
		ctx, cancel = context.WithCancel(context.Background())
		switch state {
		case Follower:
			// calculate re vote duration
			if lastState == Candidate {
				d := time.Now().Sub(now)
				if len(rf.reVoteTime) == 100 {
					rf.reVoteTime = rf.reVoteTime[1:]
				}
				rf.reVoteTime = append(rf.reVoteTime, d)
			}
			lastState = Follower
			now = time.Now()
			go rf.follower(ctx)
		case Candidate:
			lastState = Candidate
			now = time.Now()
			go rf.candidate(ctx)
		case Leader:
			if lastState == Candidate {
				d := time.Now().Sub(now)
				if len(rf.reVoteTime) == 100 {
					rf.reVoteTime = rf.reVoteTime[1:]
				}
				rf.reVoteTime = append(rf.reVoteTime, d)
			}
			lastState = Leader
			now = time.Now()
			go rf.leader(ctx)
		case Shutdown:
			cancel()
			return
		}
	}
}

// MemberChange can only change one membership at a time in case of brain split cases
func (rf *Raft) MemberChange(rpcConfig Config, md5 string) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	cOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	ok := false
	for index, newPeer := range rpcConfig.RaftPeers {
		if rf.cfg.RaftPeers[index].Port != newPeer.Port || rf.cfg.RaftPeers[index].Host != newPeer.Host {
			serverAddr := fmt.Sprintf("%s:%s", newPeer.Host, newPeer.Port)
			conn, err := grpc.Dial(serverAddr, cOpts...)
			if err != nil {
				rf.logger.Printf("open connection with id %d error, addr: %s, error: %v", newPeer.ID, serverAddr, err.Error())
				break
			}
			client := raftrpc.NewRaftClient(conn)
			rf.peers[index] = client
			if rf.state == Leader {
				rf.leaderState.nextIndex[newPeer.ID] = 0
				rf.leaderState.matchIndex[newPeer.ID] = 0
			}
			ok = true
		}
		if index == len(rpcConfig.RaftPeers)-1 {
			ok = true
		}
	}

	if ok {
		rf.cfgVersion = md5
		rf.cfg = rpcConfig
	}
}

// reportStatus add log to see instance's status
func (rf *Raft) reportStatus() {
	for {
		rf.mu.Lock()
		rf.logger.Printf("{ state: %v, term: %d, index: %d }", rf.state, rf.currentTerm, rf.currentIndex)
		rm := rf.getStatus()
		rmJson, _ := json.Marshal(rm)
		if rf.redisClient != nil {
			err := rf.redisClient.Publish(ReportChan, rmJson)
			if err != nil {
				rf.logger.Printf("publish peer status error: %v", err.Error())
			}
		}
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
					rf.logger.Printf("new entry: type=%v, entry=%v", log.Term, log.Entry)
					switch log.Type {
					case KV:
						command := string(log.Entry)
						words := strings.Split(command, " ")
						if words[0] == "set" {
							if len(words) <= 2 {
								rf.logger.Printf("entry format error kv")
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
					case LoadPrefix:
						var args LoadPrefixArgs
						if err := json.Unmarshal(log.Entry, &args); err != nil {
							rf.logs[i].Status = false
							rf.logger.Printf("parse load prefix args error: %v", err.Error())
						} else {
							rf.storage.LoadPrefix(args.Prefix, args.Data)
							rf.logger.Printf("load prefix success: %v", args.Prefix)
						}
					case RemovePrefix:
						prefix := string(log.Entry)
						rf.storage.RemovePrefix(prefix)
						rf.logger.Printf("remove prefix success: %v", prefix)
					}
					rf.lastApplied = log.Index
				}
			}
		}
		// save apply status
		rf.persist()
		rf.mu.Unlock()
		time.Sleep(NewEntryTimeout)
	}
}

func (rf *Raft) persist() {
	//rf.logger.Printf("persist raft status")
	w := new(bytes.Buffer)
	e := labgob.NewEncoder(w)
	e.Encode(rf.commitIndex)
	e.Encode(rf.lastApplied)
	e.Encode(rf.currentTerm)
	e.Encode(rf.logs)
	e.Encode(rf.votedFor)
	data := w.Bytes()
	rf.persister.SaveRaftState(data)
}

func (rf *Raft) snapshot() {
	rf.logger.Printf("persist snapshot %v\n", rf.storage.Copy())
	rf.persister.SaveSnapshot(rf.storage.Copy().(map[string]map[string]string))
	/*for i, log := range rf.logs {
		if log.Index >= rf.lastApplied {
			rf.logs = rf.logs[i:]
			break
		}
	}*/
}

func (rf *Raft) readPersist() {
	// read raft status
	data := rf.persister.ReadRaftState()
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}

	r := bytes.NewBuffer(data)
	d := labgob.NewDecoder(r)
	var commitIndex, lastApplied, currentTerm, votedFor int32
	var logs []Log
	if d.Decode(&commitIndex) != nil ||
		d.Decode(&lastApplied) != nil ||
		d.Decode(&currentTerm) != nil ||
		d.Decode(&logs) != nil ||
		d.Decode(&votedFor) != nil {
		fmt.Println("Decode Error")
	}
	rf.commitIndex, rf.lastApplied, rf.currentTerm = commitIndex, lastApplied, currentTerm
	rf.votedFor = votedFor
	rf.logs = logs

	// read snapshot
	log.Printf("read snapshot %v\n", rf.persister.ReadSnapshot())
	//rf.storage.Load(rf.persister.ReadSnapshot())
}

func (rf *Raft) getStatus() ReportMsg {
	machineMem, _ := mem.VirtualMemory()
	var curMem runtime.MemStats
	runtime.ReadMemStats(&curMem)

	var revoteTime, commitTime int64 = 0, 0
	if len(rf.reVoteTime) != 0 {
		var reVoteTimeSum time.Duration = 0
		for _, d := range rf.reVoteTime {
			reVoteTimeSum += d
		}
		revoteTime = reVoteTimeSum.Milliseconds() / int64(len(rf.reVoteTime))
	}
	if len(rf.msgCommitTime) != 0 {
		var msgCommitTimeSum time.Duration = 0
		for _, d := range rf.msgCommitTime {
			msgCommitTimeSum += d
		}
		commitTime = msgCommitTimeSum.Milliseconds() / int64(len(rf.msgCommitTime))
	}
	//rf.logger.Printf("%v\n%v\n", rf.msgCommitTime, rf.reVoteTime)
	rm := ReportMsg{
		RaftID:          rf.raftCfg.RaftID,
		Id:              int(rf.id),
		IsLeader:        rf.leaderID == rf.id,
		Status:          rf.state,
		CfgVersion:      rf.cfgVersion,
		CurrentTerm:     rf.currentTerm,
		CurrentIndex:    rf.currentIndex,
		CommitIndex:     rf.commitIndex,
		ReVoteTime:      revoteTime,
		CommitTime:      commitTime,
		MemoryTotal:     machineMem.Total,
		MemoryUsed:      machineMem.Used,
		MemoryAvailable: machineMem.Available,
		MemoryCur:       curMem.Alloc,
		Now:             time.Now(),
	}
	return rm
}

/*func (raft *Config) getConn(id int) *grpc.ClientConn {
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
