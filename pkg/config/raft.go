package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Rpc struct {
	ID   int32  `yaml:"id"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Raft struct {
	RaftRpc   Rpc    `yaml:"raft_rpc"`
	RaftPeers []Rpc  `yaml:"raft_peers"`
	LogPrefix string `yaml:"log_prefix"`
}

func NewRaftRpcConfig(path string) Raft {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("fail to read config file from %v. err: %v", path, err.Error()))
	}
	var raft Raft
	err = yaml.Unmarshal(config, &raft)
	if err != nil {
		panic(fmt.Sprintf("fail to parse config file from %v, err: %v", path, err.Error()))
	}

	return raft
}
