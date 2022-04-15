package scheduler

import (
	"configStorage/api/register"
	"configStorage/internal/raft"
	"configStorage/pkg/logger"
	"configStorage/tools/md5"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

// RegisterCenter is the center of raft addresses' registrations
type RegisterCenter struct {
	register.UnimplementedRegisterRaftServer

	register.UnimplementedKvStorageServer

	mu sync.Mutex

	raftServer *grpc.Server

	apiServer *grpc.Server

	// a simple storage interface with Get and Set functions
	s Storage

	// raftCluster
	clusters map[string]*raftCluster

	// raftIds groups names' of raft cluster
	raftIds []string

	// storage namespace
	namespace map[string]namespace

	// size raft cluster number
	size int

	// logger
	logger *logger.Logger

	cfg RegisterConfig
}

type raftCluster struct {
	once sync.Once

	// size of raft cluster
	size int

	// clusterStatus is the status of cluster membership
	status clusterStatus

	receivedCnt int

	md5 string

	// configs of raft cluster
	raftCfg []raftCfg

	// instances
	client raft.Client
}

type raftCfg struct {
	uid        string
	host       string
	clientPort string
	raftPort   string
	taken      bool
}

type namespace struct {
	raftId     string
	privateKey string
	status     bool
}

func NewRegisterCenter(config RegisterConfig) *RegisterCenter {
	return &RegisterCenter{
		s:        NewMapStorage(),
		cfg:      config,
		logger:   logger.NewLogger(make([]interface{}, 0), config.LogPrefix),
		clusters: make(map[string]*raftCluster),
		raftIds:  make([]string, 0),
		size:     0,
	}
}

func (r *RegisterCenter) Start() {
	var sOpts []grpc.ServerOption

	go func() {
		address := fmt.Sprintf("%s:%s", r.cfg.Host, r.cfg.Port)
		l, err := net.Listen("tcp", address)
		if err != nil {
			r.logger.Fatalf("Start rpc server error: %v", err.Error())
		}
		r.raftServer = grpc.NewServer(sOpts...)
		register.RegisterRegisterRaftServer(r.raftServer, r)

		r.logger.Printf("serving register center at %s:%s", r.cfg.Host, r.cfg.Port)
		err = r.raftServer.Serve(l)
		if err != nil {
			r.logger.Fatalf("Server rpc error: %v", err.Error())
		}
	}()

	address := fmt.Sprintf("%s:%s", r.cfg.Host, r.cfg.CPort)
	l, err := net.Listen("tcp", address)
	if err != nil {
		r.logger.Fatalf("Start rpc server error: %v", err.Error())
	}

	r.apiServer = grpc.NewServer(sOpts...)
	register.RegisterKvStorageServer(r.apiServer, r)

	r.logger.Printf("serving api server at %s:%s", r.cfg.Host, r.cfg.CPort)

	err = r.apiServer.Serve(l)
	if err != nil {
		r.logger.Fatalf("Server rpc error: %v", err.Error())
	}
}

func (r *RegisterCenter) RegisterRaft(ctx context.Context, args *register.RegisterRaftArgs) (reply *register.RegisterRaftReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.logger.Printf("raft register request from cluster %s, instance %s: %s:%s", args.RaftID, args.Uid, args.Host, args.RaftPort)

	reply = &register.RegisterRaftReply{OK: false}

	var cluster *raftCluster
	ok := true
	if cluster, ok = r.clusters[args.RaftID]; !ok {
		cluster = &raftCluster{
			size:        0,
			status:      Unready,
			raftCfg:     make([]raftCfg, 0),
			receivedCnt: 0,
			md5:         md5.GetRandomMd5(),
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
		taken:      true,
	}

	cluster.size++

	// fill the empty position
	if cluster.status == Changed {
		for i, _ := range cluster.raftCfg {
			if cluster.raftCfg[i].taken == false {
				cluster.raftCfg[i] = cfg
			}
		}
	} else {
		cluster.raftCfg = append(cluster.raftCfg, cfg)
	}

	if cluster.size == r.cfg.Size {
		cluster.md5 = md5.GetRandomMd5()
		cluster.receivedCnt = 0
		if cluster.status == Unready {
			cluster.status = Ready
		} else {
			cluster.status = Renew
		}
	}

	if !ok {
		r.raftIds = append(r.raftIds, args.RaftID)
		r.size++
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

	var cluster *raftCluster
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
	cluster.raftCfg[instanceId].taken = false

	if cluster.size < r.cfg.Size && cluster.status != Changed {
		cluster.status = Changed
		cluster.once = sync.Once{}
	}

	// TODO delete raft cluster size
	r.clusters[args.RaftID] = cluster

	// TODO persist cluster status
	return reply, nil
}

func (r *RegisterCenter) GetRaftRegistrations(ctx context.Context, args *register.GetRaftRegistrationsArgs) (reply *register.GetRaftRegistrationsReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reply = &register.GetRaftRegistrationsReply{OK: false, Config: make([]byte, 0)}
	cfg := raft.Config{LogPrefix: r.cfg.RaftLogPrefix}

	var cluster *raftCluster
	ok := true
	// cluster not exist
	if cluster, ok = r.clusters[args.RaftID]; !ok {
		return reply, GetDataErr
	}

	// not complete yet
	if cluster.size != r.cfg.Size {
		return reply, nil
	}

	// complete wait to connect
	go cluster.once.Do(func() {
		r.logger.Printf("trying to connect to instances......")
		r.getConn(args.RaftID)
	})

	if args.Version == cluster.md5 {
		cluster.receivedCnt++
	}

	if cluster.receivedCnt == r.cfg.Size {
		r.logger.Printf("new cfg has been received by all instances of cluster %s", args.RaftID)
		cluster.status = Ready
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
	reply.Md5 = cluster.md5
	return reply, nil
}

func (r *RegisterCenter) NewNamespace(ctx context.Context, args *register.NewNamespaceArgs) (reply *register.NewNamespaceReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.NewNamespaceReply{OK: false}

	if _, ok := r.namespace[args.Name]; ok {
		return reply, NamespaceExistedErr
	}

	name := namespace{
		raftId:     args.RaftId,
		privateKey: args.PrivateKey,
	}

	r.namespace[args.Name] = name
	reply.OK = true
	return reply, nil
}

func (r *RegisterCenter) GetClusters(ctx context.Context, args *register.GetClusterArgs) (reply *register.GetClusterReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.GetClusterReply{
		Clusters: make([]*register.GetClusterReply_Cluster, 0),
	}
	for raftID, cluster := range r.clusters {
		addr := ""
		for _, instance := range cluster.raftCfg {
			addr = addr + instance.host + ":" + instance.clientPort + "\n"
		}
		re := register.GetClusterReply_Cluster{
			RaftID:  raftID,
			Address: addr,
		}
		reply.Clusters = append(reply.Clusters, &re)
	}
	return reply, nil
}

func (r *RegisterCenter) GetConfig(ctx context.Context, args *register.GetConfigArgs) (reply *register.GetConfigReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.GetConfigReply{OK: false}
	var namespace namespace
	ok := true
	if namespace, ok = r.namespace[args.Namespace]; !ok {
		return reply, NamespaceNotExistedErr
	} else if namespace.privateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	reply.Value = r.clusters[namespace.raftId].client.Get(args.Key)
	reply.OK = true
	return reply, nil
}

func (r *RegisterCenter) GetConfigsByNamespace(ctx context.Context, args *register.GetConfigsByNamespaceArgs) (reply *register.GetConfigsByNamespaceReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.GetConfigsByNamespaceReply{OK: false}
	var namespace namespace
	ok := true
	if namespace, ok = r.namespace[args.Namespace]; !ok {
		return reply, NamespaceNotExistedErr
	} else if namespace.privateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	reply.Configs = r.clusters[namespace.raftId].client.PrefixConfig(args.Namespace)
	reply.OK = true
	return reply, nil
}

func (r *RegisterCenter) Commit(ctx context.Context, args *register.CommitArgs) (reply *register.CommitReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.CommitReply{OK: false, LastCommitID: args.Ops[0].Id}
	if r.namespace[args.Namespace].status {
		return reply, NamespaceCommittingErr
	}

	if namespace, ok := r.namespace[args.Namespace]; !ok {
		return reply, NamespaceNotExistedErr
	} else if namespace.privateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	for _, op := range args.Ops {
		if op.Type == 0 {
			if r.setConfig(args.Namespace, op) != nil {
				reply.LastCommitID = op.Id
				return reply, nil
			}
		} else if r.delConfig(args.Namespace, op) != nil {
			reply.LastCommitID = op.Id
			return reply, nil
		}
	}
	reply.OK = true
	reply.LastCommitID = args.Ops[len(args.Ops)-1].Id
	return reply, nil
}

func (r *RegisterCenter) setConfig(name string, args *register.ConfigOp) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clusters[name].client.Set(args.Key, args.Value)

	return nil
}

func (r *RegisterCenter) delConfig(name string, args *register.ConfigOp) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clusters[name].client.Del(args.Key)

	return nil
}

func (r *RegisterCenter) getConn(raftId string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	cluster := r.clusters[raftId]

	cfg := raft.ClientConfig{
		Size:      cluster.size,
		Addresses: make([]string, 0),
	}

	for _, rfCfg := range cluster.raftCfg {
		address := fmt.Sprintf("%s:%s", rfCfg.host, rfCfg.clientPort)
		cfg.Addresses = append(cfg.Addresses, address)
	}

	for times := 3; times != 0; times-- {
		time.Sleep(10 * time.Second)
		r.logger.Printf("connect with raft instances %v %v time...", cfg.Addresses, 4-times)
		if cluster.client = raft.NewRaftClient(cfg); cluster != nil {
			break
		}
	}

	r.logger.Printf("connect with raft instances success")
}
