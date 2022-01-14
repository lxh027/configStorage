package main

import (
	"configStorage/api/register"
	"configStorage/internal/raft"
	"configStorage/tools/md5"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"path"
	"time"
)

func main() {
	var env string
	var port string
	var cport string
	flag.StringVar(&env, "env", "dev", "配置环境")
	flag.StringVar(&port, "raft-port", "2001", "raft端口")
	flag.StringVar(&cport, "client-port", "3001", "raft client 端口")
	flag.Parse()

	p := path.Join("./config", env, "raft.yml")
	raftConfig := raft.NewRaftRpcConfig(p)

	uid := md5.GetRandomMd5()
	fmt.Printf("uid: %v\n", uid)
	args := register.RegisterRaftArgs{
		Uid:        uid,
		RaftID:     raftConfig.RaftID,
		Host:       raftConfig.Host,
		RaftPort:   raftConfig.Port,
		ClientPort: raftConfig.CPort,
	}

	if port != "2001" {
		args.RaftPort = port
	}

	if cport != "3001" {
		args.ClientPort = cport
	}

	cOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	registerConn, err := grpc.Dial(raftConfig.RegisterAddr, cOpts...)
	if err != nil {
		log.Fatalf("open connection error, addr: %s, error: %v", raftConfig.RegisterAddr, err.Error())
	}

	client := register.NewRegisterRaftClient(registerConn)

	_, err = client.RegisterRaft(context.Background(), &args)
	if err != nil {
		log.Fatalf("connect register center error: %v", err.Error())
	}

	log.Printf("wait for others to connect...\n")

	time.Sleep(5 * time.Second)

	reply, err := client.GetRaftRegistrations(context.Background(), &register.GetRaftRegistrationsArgs{
		Uid:    uid,
		RaftID: raftConfig.RaftID,
	})

	if err != nil {
		log.Fatalf("open connection error, addr: %s, error: %v", raftConfig.RegisterAddr, err.Error())
	}
	var cfg raft.Config
	err = json.Unmarshal(reply.Config, &cfg)

	if err != nil {
		log.Fatalf("parse config error: %v", err.Error())
	}

	// TODO defer unregister

	rafter := raft.NewRaftInstance(cfg)

	rafter.Start()
}
