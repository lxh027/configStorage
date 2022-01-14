package scheduler

import (
	"configStorage/api/register"
	"configStorage/internal/raft"
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// RegisterCenter is the center of raft addresses' registrations
type RegisterCenter struct {
	register.UnimplementedRegisterRaftServer

	mu sync.Mutex

	// a simple storage interface with Get and Set functions
	s Storage

	// raft instances' num of raft cluster
	raftSize int

	// logPrefix
	logPrefix string
}

type raftCfg struct {
	uid        string
	host       string
	clientPort string
	raftPort   string
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
	// raft cluster already full
	if peerCnt == r.raftSize {
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

func (r *RegisterCenter) GetRaftRegistrations(ctx context.Context, args *register.GetRaftRegistrationsArgs) (reply *register.GetRaftRegistrationsReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reply = &register.GetRaftRegistrationsReply{OK: false, Config: make([]byte, 0)}
	cfg := raft.Config{LogPrefix: r.logPrefix}
	cntKey := raftPeerCntKey(args.RaftID)
	peerCnt := 0
	if v, err := r.s.Get(cntKey); err != nil || v == nil {
		return reply, GetDataErr
	} else {
		peerCnt = v.(int)
	}

	// not complete yet
	if peerCnt != r.raftSize {
		return reply, nil
	}

	for idx := 0; idx < r.raftSize; idx++ {
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

func raftPeerInfoKey(id int, raftId int64) string {
	return fmt.Sprintf("raft_peer_info_%d_%d", raftId, id)
}

func raftPeerCntKey(raftId int64) string {
	return fmt.Sprintf("raft_peer_cnt_%d", raftId)
}
