package scheduler

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RegisterConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	// raft instances' num of raft cluster
	Size          int    `yaml:"size"`
	RaftLogPrefix string `yaml:"raft-log-prefix"`
	LogPrefix     string `yaml:"log-prefix"`
}

func NewRegisterConfig(path string) RegisterConfig {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("fail to read config file from %v. err: %v", path, err.Error()))
	}
	var cfg RegisterConfig
	err = yaml.Unmarshal(config, &cfg)
	if err != nil {
		panic(fmt.Sprintf("fail to parse config file from %v, err: %v", path, err.Error()))
	}
	return cfg
}
