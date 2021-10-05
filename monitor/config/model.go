package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


type Model struct {
	Save Save `yaml:"save"`
	Log  Log  `yaml:"log"`
}

type Save struct {
	Interval	uint32 	`yaml:"interval"`
	Path 		string 	`yaml:"path"`
}

type Log struct {
	Path 	string 	`yaml:"path"`
}

var modelConfig Model

func init() {
	configFile, err := ioutil.ReadFile("./config/model.yml")
	if err != nil {
		log.Printf("Read Config File Error: %v", err.Error())
		panic("Read Config File Error")
	}
	err = yaml.Unmarshal(configFile, &modelConfig)
	if err != nil {
		panic("Parse Config File Error")
	}
}

func GetModelConfig() Model {
	return modelConfig
}



