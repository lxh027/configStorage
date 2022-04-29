package scheduler

import (
	"configStorage/api/register"
	"configStorage/internal/raft"
	"configStorage/pkg/config"
	"configStorage/pkg/logger"
	"configStorage/pkg/redis"
	"configStorage/tools/md5"
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

// RegisterCenter is the center of raft addresses' registrations
type RegisterCenter struct {
	register.UnimplementedRegisterRaftServer

	register.UnimplementedKvStorageServer

	redis *redis.Client

	mu sync.Mutex

	raftServer *grpc.Server

	apiServer *grpc.Server

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
	Uid        string
	Host       string
	ClientPort string
	RaftPort   string
	Taken      bool
}

type namespace struct {
	RaftId     string
	PrivateKey string
	Status     bool
}

type persist struct {
	Clusters map[string]persistCluster
	// raftIds groups names' of raft cluster
	RaftIds []string
	// storage namespace
	Namespace map[string]namespace
	// size raft cluster number
	Size int
}

type persistCluster struct {
	Size        int
	Status      clusterStatus
	ReceivedCnt int
	Md5         string
	RaftCfg     []raftCfg
}

func NewRegisterCenter(config RegisterConfig, redisConfig config.Redis, readPersist bool) *RegisterCenter {
	redisClient, err := redis.NewRedisClient(&redisConfig)
	if err != nil {
		log.Printf("open connection with redis error: %v\n", err.Error())
		redisClient = nil
	}

	rc := &RegisterCenter{
		cfg:       config,
		redis:     redisClient,
		logger:    logger.NewLogger(make([]interface{}, 0), config.LogPrefix),
		clusters:  make(map[string]*raftCluster),
		raftIds:   make([]string, 0),
		namespace: make(map[string]namespace),
		size:      0,
	}

	if readPersist {
		if p := readPersistRc(redisClient); p != nil {
			rc.raftIds = p.RaftIds
			rc.namespace = p.Namespace
			rc.size = p.Size
			for k, c := range p.Clusters {
				rc.clusters[k] = &raftCluster{
					once:        sync.Once{},
					size:        c.Size,
					status:      c.Status,
					receivedCnt: c.ReceivedCnt,
					md5:         c.Md5,
					raftCfg:     c.RaftCfg,
				}
				rc.getConn(k)
			}
			log.Printf("read persist success: %v", p)
		}
	}
	return rc
}

func (r *RegisterCenter) Start() {
	sOpts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger.NewZapLogger()))),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger.NewZapLogger()))),
	}

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

	for i, instance := range cluster.raftCfg {
		if instance.Host == args.Host && instance.ClientPort == args.ClientPort {
			cluster.size--
			cluster.raftCfg[i].Taken = false
			if cluster.size < r.cfg.Size && cluster.status != Changed {
				cluster.status = Changed
				cluster.once = sync.Once{}
			}
		}
	}

	// check if raft cluster is full
	if cluster.size == r.cfg.Size {
		return reply, RaftFullErr
	}

	// check if uid is in cluster
	for _, instance := range cluster.raftCfg {
		if instance.Uid == args.Uid {
			return reply, RaftInstanceExistedErr
		}
	}

	cfg := raftCfg{
		Uid:        args.Uid,
		Host:       args.Host,
		ClientPort: args.ClientPort,
		RaftPort:   args.RaftPort,
		Taken:      true,
	}

	cluster.size++

	// fill the empty position
	if cluster.status == Changed {
		for i, _ := range cluster.raftCfg {
			if cluster.raftCfg[i].Taken == false {
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
		if instance.Uid == args.Uid {
			instanceId = id
			break
		}
	}

	if instanceId == -1 {
		return reply, RaftInstanceNotExistedErr
	}

	cluster.size--
	cluster.raftCfg[instanceId].Taken = false

	if cluster.size < r.cfg.Size && cluster.status != Changed {
		cluster.status = Changed
		cluster.once = sync.Once{}
	}

	r.clusters[args.RaftID] = cluster

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
		rpc := raft.Rpc{ID: int32(idx), Host: instance.Host, Port: instance.RaftPort, CPort: instance.ClientPort}
		cfg.RaftPeers = append(cfg.RaftPeers, rpc)
		if instance.Uid == args.Uid {
			cfg.RaftRpc = rpc
		}
	}
	// persist register
	persistRc(r)
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
		RaftId:     args.RaftId,
		PrivateKey: args.PrivateKey,
	}

	r.namespace[args.Name] = name
	reply.OK = true
	// persist register
	persistRc(r)
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
			addr = addr + instance.Host + ":" + instance.ClientPort + "\n"
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
	} else if namespace.PrivateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	redisKey := fmt.Sprintf("configCache.%s.%s", args.Namespace, args.Key)
	if r.redis != nil {
		if value, err := r.redis.GetFromRedis(redisKey); err == nil {
			reply.Value = value.(string)
			reply.OK = true
			return reply, nil
		}
	}

	reply.Value = r.clusters[namespace.RaftId].client.Get(args.Key)
	if r.redis != nil {
		_ = r.redis.PutToRedis(redisKey, reply.Value, 300)
	}
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
	} else if namespace.PrivateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	reply.Configs = r.clusters[namespace.RaftId].client.PrefixConfig(args.Namespace)
	r.logger.Printf("config: %v", reply.Configs)
	reply.OK = true
	return reply, nil
}

func (r *RegisterCenter) Commit(ctx context.Context, args *register.CommitArgs) (reply *register.CommitReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.logger.Printf("start committing...namespace: %v, ops: %v", args.Namespace, args.Ops)
	reply = &register.CommitReply{OK: false, LastCommitID: args.Ops[0].Id}

	if len(args.Ops) == 0 {
		return reply, nil
	}

	if namespace, ok := r.namespace[args.Namespace]; !ok {
		return reply, NamespaceNotExistedErr
	} else if namespace.PrivateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	for _, op := range args.Ops {
		go r.delRedis(args.Namespace, op.Key)
		if op.Type == 0 {
			if err := r.setConfig(args.Namespace, op); err != nil {
				reply.LastCommitID = op.Id
				return reply, err
			}
		} else if err := r.delConfig(args.Namespace, op); err != nil {
			reply.LastCommitID = op.Id
			return reply, err
		}
	}
	reply.OK = true
	reply.LastCommitID = args.Ops[len(args.Ops)-1].Id
	return reply, nil
}

func (r *RegisterCenter) DeleteNamespace(ctx context.Context, args *register.DeleteNamespaceArgs) (reply *register.DeleteNamespaceReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.DeleteNamespaceReply{}
	if namespace, ok := r.namespace[args.Namespace]; !ok {
		return reply, NamespaceNotExistedErr
	} else if namespace.PrivateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	err = r.clusters[r.namespace[args.Namespace].RaftId].client.PrefixRemove(args.Namespace)
	if err != nil {
		return reply, err
	}
	delete(r.namespace, args.Namespace)
	// persist register
	persistRc(r)
	return reply, nil
}

func (r *RegisterCenter) TransNamespace(ctx context.Context, args *register.TransNamespaceArgs) (reply *register.TransNamespaceReply, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reply = &register.TransNamespaceReply{}

	var namespace namespace
	var ok bool
	if namespace, ok = r.namespace[args.Namespace]; !ok {
		return reply, NamespaceNotExistedErr
	} else if namespace.PrivateKey != args.PrivateKey {
		return reply, PrivateKeyUnPatchErr
	}

	if _, ok := r.clusters[args.RaftID]; !ok {
		return reply, ClusterNotExistedErr
	}

	configs := r.clusters[namespace.RaftId].client.PrefixConfig(args.Namespace)
	err = r.clusters[args.RaftID].client.PrefixLoad(args.Namespace, configs)
	if err != nil {
		return reply, err
	}
	err = r.clusters[r.namespace[args.Namespace].RaftId].client.PrefixRemove(args.Namespace)
	if err != nil {
		return reply, err
	}
	namespace.RaftId = args.RaftID
	r.namespace[args.Namespace] = namespace
	return reply, nil
}

func (r *RegisterCenter) setConfig(name string, args *register.ConfigOp) error {
	raftID := r.namespace[name].RaftId
	return r.clusters[raftID].client.Set(name+"."+args.Key, args.Value)
}

func (r *RegisterCenter) delConfig(name string, args *register.ConfigOp) error {
	raftID := r.namespace[name].RaftId
	return r.clusters[raftID].client.Del(name + "." + args.Key)
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
		address := fmt.Sprintf("%s:%s", rfCfg.Host, rfCfg.ClientPort)
		cfg.Addresses = append(cfg.Addresses, address)
	}
	r.logger.Printf("client cfg: %v", cfg)

	for times := 3; times != 0; times-- {
		time.Sleep(10 * time.Second)
		r.logger.Printf("connect with raft instances %v %v time...", cfg.Addresses, 4-times)
		if cluster.client = raft.NewRaftClient(cfg); cluster != nil {
			break
		}
	}

	r.logger.Printf("connect with raft instances success")
}

func (r *RegisterCenter) delRedis(namespace, key string) {
	if r.redis == nil {
		return
	}
	redisKey := fmt.Sprintf("configCache.%s.%s", namespace, key)
	_ = r.redis.DeleteFromRedis(redisKey)
}

func persistRc(r *RegisterCenter) {
	if r.redis == nil {
		return
	}
	p := persist{
		Size:      r.size,
		Namespace: r.namespace,
		RaftIds:   r.raftIds,
		Clusters:  make(map[string]persistCluster),
	}
	for k, c := range r.clusters {
		p.Clusters[k] = persistCluster{
			Size:        c.size,
			Status:      c.status,
			ReceivedCnt: c.receivedCnt,
			Md5:         c.md5,
			RaftCfg:     c.raftCfg,
		}
	}
	if bts, err := json.Marshal(p); err == nil {
		err = r.redis.PutToRedisLast("register.persist_status", bts)
		if err == nil {
			log.Printf("persist success: %v\n", p)
		}
	}
}

func readPersistRc(client *redis.Client) *persist {
	if client == nil {
		return nil
	}
	var p persist
	if bts, err := client.GetFromRedis("register.persist_status"); err == nil && bts != nil {
		err = json.Unmarshal(bts.([]byte), &p)
		if err == nil {
			log.Printf("read persist success: %v\n", p)
			return &p
		} else {
			log.Printf("read persist error: %v", err.Error())
		}
	}
	return nil
}
