package scheduler

import (
	"configStorage/api/register"
	"context"
	"google.golang.org/grpc"
)

type Log struct {
	ID    int
	Type  int
	Key   string
	Value string
}

type Cluster struct {
	RaftID  string
	Address string
}

type Client interface {
	NewNamespace(name string, privateKey string, raftID string) error
	GetClusters() ([]Cluster, error)
	GetConfig(namespace string, privateKey string, key string) (string, error)
	GetConfigByNamespace(namespace string, privateKey string) (map[string]string, error)
	Commit(namespace string, privateKey string, configs []Log) (int, error)
}

type SCDClient struct {
	register.KvStorageClient
}

func NewSchedulerClient(address string) Client {
	c := SCDClient{}
	cOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(address, cOpts...)
	if err != nil {
		return nil
	}
	c.KvStorageClient = register.NewKvStorageClient(conn)
	return &c
}

func (s *SCDClient) NewNamespace(name string, privateKey string, raftID string) error {
	args := register.NewNamespaceArgs{
		Name:       name,
		PrivateKey: privateKey,
		RaftId:     raftID,
	}
	_, err := s.KvStorageClient.NewNamespace(context.Background(), &args)
	return err
}

func (s *SCDClient) GetClusters() ([]Cluster, error) {
	reply, err := s.KvStorageClient.GetClusters(context.Background(), &register.GetClusterArgs{})
	if err != nil {
		return nil, err
	}
	clusters := make([]Cluster, 0)
	for _, c := range reply.Clusters {
		clusters = append(clusters, Cluster{RaftID: c.RaftID, Address: c.Address})
	}
	return clusters, nil
}

func (s *SCDClient) GetConfig(namespace string, privateKey string, key string) (string, error) {
	args := register.GetConfigArgs{
		Namespace:  namespace,
		Key:        key,
		PrivateKey: privateKey,
	}

	reply, err := s.KvStorageClient.GetConfig(context.Background(), &args)
	return reply.Value, err
}

func (s *SCDClient) GetConfigByNamespace(namespace string, privateKey string) (map[string]string, error) {
	args := register.GetConfigsByNamespaceArgs{
		Namespace:  namespace,
		PrivateKey: privateKey,
	}
	reply, err := s.KvStorageClient.GetConfigsByNamespace(context.Background(), &args)
	return reply.Configs, err
}

func (s *SCDClient) Commit(namespace string, privateKey string, configs []Log) (int, error) {
	ops := make([]*register.ConfigOp, 0)
	for _, c := range configs {
		ops = append(ops, &register.ConfigOp{Id: int64(c.ID), Key: c.Key, Value: c.Value, Type: int64(c.Type)})
	}

	args := register.CommitArgs{
		Namespace:  namespace,
		PrivateKey: privateKey,
		Ops:        ops,
	}

	reply, err := s.KvStorageClient.Commit(context.Background(), &args)
	return int(reply.LastCommitID), err
}