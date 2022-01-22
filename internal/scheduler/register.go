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

	// raftCluster
	clusters map[string]raftCluster

	// raftIds groups names' of raft cluster
	raftIds []string

	// logger
	logger *logger.Logger

	cfg RegisterConfig
}

type raftCluster struct {
	// size of raft cluster
	size int

	// configs of raft cluster
	raftCfg []raftCfg

	// instances
	instance map[string]raft.Client
}

type raftCfg struct {
	uid        string
	host       string
	clientPort string
	raftPort   string
}

func NewRegisterCenter(config RegisterConfig) *RegisterCenter {
	return &RegisterCenter{
		s:        NewMapStorage(),
		cfg:      config,
		logger:   logger.NewLogger(make([]interface{}, 0), config.LogPrefix),
		clusters: make(map[string]raftCluster),
		raftIds:  make([]string, 0),
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

	var cluster raftCluster
	ok := true
	if cluster, ok = r.clusters[args.RaftID]; !ok {
		cluster = raftCluster{
			size:     0,
			raftCfg:  make([]raftCfg, 0),
			instance: make(map[string]raft.Client),
		}
	}

	r.logger.Printf("peer count: %v, target count: %v", cluster.size, r.cfg.Size)

	// check if raft cluster is full
	if cluster.size == r.cfg.Size {
		return reply, RaftFullErr
	}

	// check if uid is in cluster
	for _, instance := range cluster.raftCfg {
		if instance.uid == args.Uid {
			return reply, RaftInstanceExistedErr
		}
	}

	cfg := raftCfg{
		uid:        args.Uid,
		host:       args.Host,
		clientPort: args.ClientPort,
		raftPort:   args.RaftPort,
	}

	cluster.size++
	cluster.raftCfg = append(cluster.raftCfg, cfg)

	if !ok {
		r.raftIds = append(r.raftIds, args.RaftID)
	}
	r.clusters[args.RaftID] = cluster

	// TODO persist cluster status
	reply.OK = true
	return reply, nil
}

func (r *RegisterCenter) UnregisterRaft(ctx context.Context, args *register.UnregisterRaftArgs) (reply *register.UnregisterRaftReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reply = &register.UnregisterRaftReply{}

	var cluster raftCluster
	ok := true
	// cluster not exist
	if cluster, ok = r.clusters[args.RaftID]; !ok {
		return reply, GetDataErr
	}

	r.logger.Printf("peer count: %v, target count: %v", cluster.size, r.cfg.Size)

	// check if raft cluster is empty
	if cluster.size <= 0 {
		return reply, RaftEmptyErr
	}

	// check if raft id is existed
	instanceId := -1
	for id, instance := range cluster.raftCfg {
		if instance.uid == args.Uid {
			instanceId = id
			break
		}
	}

	if instanceId == -1 {
		return reply, RaftInstanceNotExistedErr
	}

	cluster.size--
	cluster.raftCfg = append(cluster.raftCfg[:instanceId], cluster.raftCfg[instanceId+1:]...)

	delete(cluster.instance, args.Uid)

	r.clusters[args.RaftID] = cluster

	// TODO persist cluster status
	return reply, nil
}

func (r *RegisterCenter) GetRaftRegistrations(ctx context.Context, args *register.GetRaftRegistrationsArgs) (reply *register.GetRaftRegistrationsReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reply = &register.GetRaftRegistrationsReply{OK: false, Config: make([]byte, 0)}
	cfg := raft.Config{LogPrefix: r.cfg.RaftLogPrefix}

	var cluster raftCluster
	ok := true
	// cluster not exist
	if cluster, ok = r.clusters[args.RaftID]; !ok {
		return reply, GetDataErr
	}

	// not complete yet
	if cluster.size != r.cfg.Size {
		return reply, nil
	}

	for idx, instance := range cluster.raftCfg {
		rpc := raft.Rpc{ID: int32(idx), Host: instance.host, Port: instance.raftPort, CPort: instance.clientPort}
		cfg.RaftPeers = append(cfg.RaftPeers, rpc)
		if instance.uid == args.Uid {
			cfg.RaftRpc = rpc
		}
	}

	byteData, _ := json.Marshal(cfg)
	reply.OK = true
	reply.Config = byteData
	return reply, nil
}
