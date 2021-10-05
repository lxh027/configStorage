package model

import (
	"errors"
	"log"
	"monitor/model/enum"
	"monitor/model/models"
	"monitor/model/storage"
)

var modelStorage storage.ModelStorage


func StartStorage() {
	modelStorage.Init()
}

func AddModel(modelType enum.ModelType, instanceId uint32, title string) {
	switch modelType {
	case enum.Log:
		logModel := models.NewLogModel(instanceId)
		modelStorage.Models = append(modelStorage.Models, logModel)
	default:
		log.Println("undefined log type")
	}
}

func GetAllModels() []models.Model {
	return modelStorage.Models
}

func GetSingleModel(modelType enum.ModelType, instanceID uint32) (models.Model, error)  {
	for _, model := range modelStorage.Models {
		if model.GetParams().Check(modelType, instanceID) {
			return model, nil
		}
	}
	return nil, errors.New("find model error")
}
