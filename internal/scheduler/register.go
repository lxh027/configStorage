package scheduler

import (
	"configStorage/api/register"
	"configStorage/internal/raft"
	"configStorage/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
)

// RegisterCenter is the center of raft addresses' registrations
type RegisterCenter struct {
	register.UnimplementedRegisterRaftServer

	mu sync.Mutex

	server *grpc.Server

	// a simple storage interface with Get and Set functions
	s Storage

	// logger
	logger *logger.Logger

	cfg RegisterConfig
}

type raftCfg struct {
	uid        string
	host       string
	clientPort string
	raftPort   string
}

func NewRegisterCenter(config RegisterConfig) *RegisterCenter {
	return &RegisterCenter{
		s:      NewMapStorage(),
		cfg:    config,
		logger: logger.NewLogger(make([]interface{}, 0), config.LogPrefix),
	}
}

func (r *RegisterCenter) Start() {
	address := fmt.Sprintf("%s:%s", r.cfg.Host, r.cfg.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		r.logger.Fatalf("Start rpc server error: %v", err.Error())
	}
	var sOpts []grpc.ServerOption

	r.server = grpc.NewServer(sOpts...)
	register.RegisterRegisterRaftServer(r.server, r)

	r.logger.Printf("serving register center at %s:%s", r.cfg.Host, r.cfg.Port)

	err = r.server.Serve(l)
	if err != nil {
		r.logger.Fatalf("Server rpc error: %v", err.Error())
	}
}

func (r *RegisterCenter) RegisterRaft(ctx context.Context, args *register.RegisterRaftArgs) (reply *register.RegisterRaftReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.RegisterRaftReply{OK: false}
	peerCnt := 0
	cntKey := raftPeerCntKey(args.RaftID)
	v, err := r.s.Get(cntKey)
	if err != nil {
		return reply, GetDataErr
	}
	if v != nil {
		peerCnt = v.(int)
	}
	r.logger.Printf("peer count: %v, target count: %v", peerCnt, r.cfg.Size)
	// raft cluster already full
	if peerCnt == r.cfg.Size {
		return reply, RaftFullErr
	}
	// new raft peer

	cfg := raftCfg{
		uid:        args.Uid,
		host:       args.Host,
		clientPort: args.ClientPort,
		raftPort:   args.RaftPort,
	}

	key := raftPeerInfoKey(peerCnt, args.RaftID)

	if r.s.Set(key, cfg) != nil {
		return reply, SetDataErr
	}

	peerCnt++
	_ = r.s.Set(cntKey, peerCnt)
	reply.OK = true
	return reply, nil
}

func (r *RegisterCenter) UnregisterRaft(ctx context.Context, args *register.UnregisterRaftArgs) (reply *register.UnregisterRaftReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reply = &register.UnregisterRaftReply{}

	cntKey := raftPeerCntKey(args.RaftID)
	key := raftPeerInfoKey(int(args.Idx), args.RaftID)
	peerCnt := 0

	v, err := r.s.Get(cntKey)
	if err != nil {
		return reply, GetDataErr
	}
	if v != nil {
		peerCnt = v.(int)
	}
	r.logger.Printf("peer count: %v, target count: %v", peerCnt, r.cfg.Size)

	// check if raft cluster is empty
	if peerCnt <= 0 {
		return reply, RaftEmptyErr
	}

	if r.s.Del(key) != nil {
		return reply, DelDataErr
	}

	peerCnt--

	_ = r.s.Set(cntKey, peerCnt)
	return reply, nil
}

func (r *RegisterCenter) GetRaftRegistrations(ctx context.Context, args *register.GetRaftRegistrationsArgs) (reply *register.GetRaftRegistrationsReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reply = &register.GetRaftRegistrationsReply{OK: false, Config: make([]byte, 0)}
	cfg := raft.Config{LogPrefix: r.cfg.RaftLogPrefix}
	cntKey := raftPeerCntKey(args.RaftID)
	peerCnt := 0
	if v, err := r.s.Get(cntKey); err != nil || v == nil {
		return reply, GetDataErr
	} else {
		peerCnt = v.(int)
	}

	// not complete yet
	if peerCnt != r.cfg.Size {
		return reply, nil
	}

	for idx := 0; idx < r.cfg.Size; idx++ {
		key := raftPeerInfoKey(idx, args.RaftID)
		if v, err := r.s.Get(key); err != nil || v == nil {
			return reply, GetDataErr
		} else {
			peerCfg := v.(raftCfg)
			rpc := raft.Rpc{ID: int32(idx), Host: peerCfg.host, Port: peerCfg.raftPort, CPort: peerCfg.clientPort}
			cfg.RaftPeers = append(cfg.RaftPeers, rpc)
			if peerCfg.uid == args.Uid {
				cfg.RaftRpc = rpc
			}
		}
	}

	byteData, _ := json.Marshal(cfg)
	reply.OK = true
	reply.Config = byteData
	return reply, nil
}

func raftPeerInfoKey(id int, raftId string) string {
	return fmt.Sprintf("raft_peer_info_%s_%d", raftId, id)
}

func raftPeerCntKey(raftId string) string {
	return fmt.Sprintf("raft_peer_cnt_%s", raftId)
}
