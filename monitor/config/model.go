package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	configFile, err := ioutil.ReadFile("./model.yml")
	if err != nil {
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



