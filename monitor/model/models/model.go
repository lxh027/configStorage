package models

import "monitor/model/enum"

// Params basic params for any model
type Params struct {
	ModelType 		enum.ModelType	`json:"model_type"`
	InstanceID 		uint32			`json:"instance_id"`
	Title 			string			`json:"title"`
}

// Model interface defined of models
type Model interface {
	Update(param uint8, data interface{}) error
	Save(path string) error
	This() interface{}
	GetParams() *Params
	GetData() 	interface{}
}

// Check check if model match the wanted model
func (params *Params) Check(modelType enum.ModelType, instanceId uint32) bool {
	return params.ModelType == modelType && params.InstanceID == instanceId
}
