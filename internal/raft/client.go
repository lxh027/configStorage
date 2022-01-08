package raft

import (
	"configStorage/api/raftrpc"
	"configStorage/pkg/logger"
	"configStorage/tools/random"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type ClientConfig struct {
	size      int
	addresses []string
}

type Client interface {
	Get(string) string
	Set(string, string)
	Del(string)
}

type rfClient struct {
	size      int
	logger    *logger.Logger
	instances []raftrpc.StateClient
	leaderId  int
}

func NewRaftClient(cfg ClientConfig) Client {
	c := rfClient{
		size:      cfg.size,
		logger:    logger.NewLogger(make([]interface{}, 0), ""),
		instances: make([]raftrpc.StateClient, cfg.size),
		leaderId:  0,
	}

	for i, address := range cfg.addresses {
		cOpts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		conn, err := grpc.Dial(address, cOpts...)
		if err != nil {
			c.logger.Fatalf("error while get conn for raft server, error: %v", err.Error())
		}
		c.instances[i] = raftrpc.NewStateClient(conn)
	}

	return &c
}

func (c *rfClient) Set(key string, value string) {
	command := fmt.Sprintf("set %s %s", key, value)
	entry := raftrpc.NewEntryArgs{Entry: []byte(command)}

	reply, _ := c.instances[c.leaderId].NewEntry(context.Background(), &entry)
	if reply.Success {
		c.logger.Printf("set kv ok: %v %v", key, value)
		return
	}

	if int(reply.LeaderID) != c.leaderId {
		c.leaderId = int(reply.LeaderID)
		c.logger.Printf("leader changed, resent")
		reply, _ = c.instances[c.leaderId].NewEntry(context.Background(), &entry)
		if reply.Success {
			c.logger.Printf("set kv ok: %v %v", key, value)
			return
		}
	}
	c.logger.Printf("set kv error: %v %v, err: %v", key, value, reply.Msg)
}

func (c *rfClient) Get(key string) string {
	index := random.ID(c.size)
	reply, _ := c.instances[index].GetValue(context.Background(), &raftrpc.GetValueArgs{Key: key})
	c.logger.Printf("get value of key [%v: %v]", key, reply.Value)
	return reply.Value
}

func (c *rfClient) Del(key string) {
	command := fmt.Sprintf("del %s", key)
	entry := raftrpc.NewEntryArgs{Entry: []byte(command)}

	reply, _ := c.instances[c.leaderId].NewEntry(context.Background(), &entry)
	if reply.Success {
		c.logger.Printf("del kv ok: %v", key)
		return
	}

	if int(reply.LeaderID) != c.leaderId {
		c.leaderId = int(reply.LeaderID)
		c.logger.Printf("leader changed, resent")
		reply, _ = c.instances[c.leaderId].NewEntry(context.Background(), &entry)
		if reply.Success {
			c.logger.Printf("del kv ok: %v", key)
			return
		}
	}
	c.logger.Printf("del kv error: %v, err: %v", key, reply.Msg)
}
