package raft

import (
	"configStorage/api/raftrpc"
	"configStorage/pkg/logger"
	"configStorage/tools/random"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
)

type ClientConfig struct {
	Size      int
	Addresses []string
}

type Client interface {
	Get(string) string
	Set(string, string) error
	Del(string) error

	PrefixLoad(string, map[string]string) error
	PrefixRemove(string) error
	PrefixConfig(string) map[string]string
}

type rfClient struct {
	size      int
	logger    *logger.Logger
	instances []raftrpc.StateClient
	leaderId  int
}

func NewRaftClient(cfg ClientConfig) Client {
	c := rfClient{
		size:      cfg.Size,
		logger:    logger.NewLogger(make([]interface{}, 0), ""),
		instances: make([]raftrpc.StateClient, cfg.Size),
		leaderId:  0,
	}

	for i, address := range cfg.Addresses {
		cOpts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		conn, err := grpc.Dial(address, cOpts...)
		if err != nil {
			c.logger.Printf("error while get conn for raft server, error: %v", err.Error())
			return nil
		}
		c.instances[i] = raftrpc.NewStateClient(conn)
	}

	return &c
}

func (c *rfClient) PrefixLoad(prefix string, data map[string]string) error {
	loadArgs := LoadPrefixArgs{
		Prefix: prefix,
		Data:   data,
	}
	bts, err := json.Marshal(loadArgs)
	if err != nil {
		c.logger.Printf("parse data error %v", err)
		return err
	}
	entry := raftrpc.NewEntryArgs{Entry: bts, Type: int32(LoadPrefix)}

	reply, err := c.instances[c.leaderId].NewEntry(context.Background(), &entry)
	if err != nil {
		return err
	}
	if reply.Success {
		c.logger.Printf("load prefix ok: %v ", prefix)
		return nil
	}

	if int(reply.LeaderID) != c.leaderId {
		c.leaderId = int(reply.LeaderID)
		c.logger.Printf("leader changed, resent")
		reply, err = c.instances[c.leaderId].NewEntry(context.Background(), &entry)
		if err != nil {
			return err
		}
		if reply.Success {
			c.logger.Printf("load prefix ok: %v ", prefix)
			return nil
		}
	}
	c.logger.Printf("load prefix ok: %v, err: %v", prefix, reply.Msg)
	return errors.New(reply.Msg)
}

func (c *rfClient) PrefixRemove(prefix string) error {
	entry := raftrpc.NewEntryArgs{Entry: []byte(prefix), Type: int32(RemovePrefix)}

	reply, err := c.instances[c.leaderId].NewEntry(context.Background(), &entry)
	if err != nil {
		return err
	}
	if reply.Success {
		c.logger.Printf("remove prefix ok: %v ", prefix)
		return nil
	}

	if int(reply.LeaderID) != c.leaderId {
		c.leaderId = int(reply.LeaderID)
		c.logger.Printf("leader changed, resent")
		reply, err = c.instances[c.leaderId].NewEntry(context.Background(), &entry)
		if err != nil {
			return err
		}
		if reply.Success {
			c.logger.Printf("remove prefix ok: %v ", prefix)
			return nil
		}
	}
	c.logger.Printf("remove prefix ok: %v, err: %v", prefix, reply.Msg)
	return errors.New(reply.Msg)
}

func (c *rfClient) PrefixConfig(prefix string) map[string]string {
	index := random.ID(c.size)
	reply, _ := c.instances[index].GetPrefixConfigs(context.Background(), &raftrpc.GetPrefixConfigArgs{Prefix: prefix})
	return reply.Config
}

func (c *rfClient) Get(key string) string {
	index := random.ID(c.size)
	reply, _ := c.instances[index].GetValue(context.Background(), &raftrpc.GetValueArgs{Key: key})
	c.logger.Printf("get value of key [%v: %v]", key, reply.Value)
	return reply.Value
}

func (c *rfClient) Set(key string, value string) error {
	command := fmt.Sprintf("set %s %s", key, value)
	entry := raftrpc.NewEntryArgs{Entry: []byte(command), Type: int32(KV)}

	reply, err := c.instances[c.leaderId].NewEntry(context.Background(), &entry)
	if err != nil {
		return err
	}
	if reply.Success {
		c.logger.Printf("set kv ok: %v %v", key, value)
		return nil
	}

	if int(reply.LeaderID) != c.leaderId {
		c.leaderId = int(reply.LeaderID)
		c.logger.Printf("leader changed, resent")
		reply, err = c.instances[c.leaderId].NewEntry(context.Background(), &entry)
		if err != nil {
			return err
		}
		if reply.Success {
			c.logger.Printf("set kv ok: %v %v", key, value)
			return nil
		}
	}
	c.logger.Printf("set kv error: %v %v, err: %v", key, value, reply.Msg)
	return nil
}

func (c *rfClient) Del(key string) error {
	command := fmt.Sprintf("del %s", key)
	entry := raftrpc.NewEntryArgs{Entry: []byte(command), Type: int32(KV)}

	reply, err := c.instances[c.leaderId].NewEntry(context.Background(), &entry)
	if err != nil {
		return err
	}
	if reply.Success {
		c.logger.Printf("del kv ok: %v", key)
		return nil
	}

	if int(reply.LeaderID) != c.leaderId {
		c.leaderId = int(reply.LeaderID)
		c.logger.Printf("leader changed, resent")
		reply, err = c.instances[c.leaderId].NewEntry(context.Background(), &entry)
		if err != nil {
			return err
		}
		if reply.Success {
			c.logger.Printf("del kv ok: %v", key)
			return nil
		}
	}
	c.logger.Printf("del kv error: %v, err: %v", key, reply.Msg)
	return nil
}
