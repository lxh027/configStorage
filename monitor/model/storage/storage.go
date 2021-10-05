package storage

import (
	"fmt"
	"log"
	"monitor/config"
	"monitor/model/enum"
	"monitor/model/models"
	"time"
)

// struct to store models and configs
type ModelStorage struct {
	Models 			[]models.Model
	// interval to save models
	saveInterval 	time.Duration
	// path prefix of models
	modelPath 		string
	// path prefix of logs
	logPath 		string
}

// init storage
func (m *ModelStorage) Init() {
	m.Models = make([]models.Model, 0)
	m.saveInterval = time.Duration(config.GetModelConfig().Save.Interval)
	m.modelPath = config.GetModelConfig().Save.Path
	m.logPath = config.GetModelConfig().Log.Path

	// start the goroutine to save models
	go m.save()
}

// periodically save models
func (m *ModelStorage) save() {
	for {
		for _, model := range m.Models {
			params := model.GetParams()
			var path string
			// generate save path
			if params.ModelType == enum.Log {
				path = m.logPath
			} else {
				path = m.modelPath
			}
			// file patten: ${model_type}_${instance_id}
			path = path + fmt.Sprintf("/%d_%d", params.ModelType, params.InstanceID)
			err := model.Save(path)
			if err != nil {
				log.Printf("save model error: %v\n", err.Error())
			}
		}
		time.Sleep(m.saveInterval)
	}
}



