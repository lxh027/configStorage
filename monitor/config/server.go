package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Server struct {
	Grpc 	Grpc 	`yaml:"grpc"`
}

type Grpc struct {
	Port 	uint32 	`yaml:"port"`
	Host 	string 	`yaml:"host"`
}


var serverConfig Server

func init() {
	configFile, err := ioutil.ReadFile("./config/server.yml")
	if err != nil {
		log.Printf("Read Config File Error: %v", err.Error())
		panic("Read Config File Error")
	}
	err = yaml.Unmarshal(configFile, &serverConfig)
	if err != nil {
		panic("Parse Config File Error")
	}
}

func GetServerConfig() Server {
	return serverConfig
}

