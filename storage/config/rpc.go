package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Rpc struct {
	ID   int32  `yaml:"id"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type RaftRpc Rpc
type PeerRpc Rpc

type MonitorRpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type RpcConfig struct {
	RaftRpc    RaftRpc    `yaml:"raft_rpc"`
	RaftPeers  []PeerRpc  `yaml:"raft_peers"`
	MonitorRpc MonitorRpc `yaml:"monitor_rpc"`
}

var rpcConfig RpcConfig

func init() {
	configFile, err := ioutil.ReadFile("./config/rpc.yml")
	if err != nil {
		log.Printf("Read Config File Error: %v\n", err.Error())
		panic("Read config file error")
	}
	err = yaml.Unmarshal(configFile, &rpcConfig)
	if err != nil {
		panic("Parse config file error")
	}
}

func GetRpcConfig() RpcConfig {
	return rpcConfig
}
