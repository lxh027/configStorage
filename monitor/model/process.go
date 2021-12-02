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

// AddModel add new model to storage
// switch modelType to decide witch model to generate
func AddModel(modelType enum.ModelType, instanceId uint32) {
	switch modelType {
	case enum.Log:
		logModel := models.NewLogModel(instanceId)
		modelStorage.Models = append(modelStorage.Models, logModel)
	case enum.Entry:
		entryModel := models.NewEntryModel(instanceId)
		modelStorage.Models = append(modelStorage.Models, entryModel)
	default:
		log.Println("undefined log type")
	}
}

func UpdateModel(modelType enum.ModelType, instanceId uint32, param uint32, data interface{}) error {
	for index, model := range modelStorage.Models {
		if model.GetParams().Check(modelType, instanceId) {
			return modelStorage.Models[index].Update(param, data)
		}
	}
	return errors.New("model not found")
}

// GetAllModels return all models
func GetAllModels() []models.Model {
	return modelStorage.Models
}

// GetSingleModel return single model by modelType and instanceID
func GetSingleModel(modelType enum.ModelType, instanceID uint32) (models.Model, error) {
	for _, model := range modelStorage.Models {
		if model.GetParams().Check(modelType, instanceID) {
			return model, nil
		}
	}
	return nil, errors.New("find model error")
}
