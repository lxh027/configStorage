package storage

import (
	"fmt"
	"log"
	"monitor/config"
	"monitor/model/enum"
	"monitor/model/models"
	"time"
)

type ModelStorage struct {
	Models 			[]models.Model
	saveInterval 	time.Duration
	modelPath 		string
	logPath 		string
}

func (m *ModelStorage) Init() {
	m.Models = make([]models.Model, 0)
	m.saveInterval = time.Duration(config.GetModelConfig().Save.Interval)
	m.modelPath = config.GetModelConfig().Save.Path
	m.logPath = config.GetModelConfig().Log.Path

	go m.save()
}

func (m *ModelStorage) save() {
	for {
		for _, model := range m.Models {
			params := model.GetParams()
			var path string
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



