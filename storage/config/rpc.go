package config

type Rpc struct {
	ID   int32  `yaml:"id"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MonitorRpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type RpcConfig struct {
	RaftRpc    Rpc        `yaml:"raft_rpc"`
	RaftPeers  []Rpc      `yaml:"raft_peers"`
	MonitorRpc MonitorRpc `yaml:"monitor_rpc"`
}
