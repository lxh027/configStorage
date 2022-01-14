package raft

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Rpc struct {
	ID    int32
	Host  string
	Port  string
	CPort string
}

type Config struct {
	RaftRpc   Rpc
	RaftPeers []Rpc
	LogPrefix string
}

type RfConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	CPort        string `yaml:"c-port"`
	RegisterAddr string `yaml:"register-address"`
	RaftID       string `yaml:"raft-id"`
}

func NewRaftRpcConfig(path string) RfConfig {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("fail to read config file from %v. err: %v", path, err.Error()))
	}
	var raft RfConfig
	err = yaml.Unmarshal(config, &raft)
	if err != nil {
		panic(fmt.Sprintf("fail to parse config file from %v, err: %v", path, err.Error()))
	}

	return raft
}
