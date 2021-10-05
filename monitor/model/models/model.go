package models

import "monitor/model/enum"

type Params struct {
	ModelType 		enum.ModelType	`json:"model_type"`
	InstanceID 		uint32			`json:"instance_id"`
	Title 			string			`json:"title"`
}

type Model interface {
	Update(param uint8, data interface{}) error
	Save(path string) error
	This() interface{}
	GetParams() *Params
	GetData() 	interface{}
}

func (params *Params) Check(modelType enum.ModelType, instanceId uint32) bool {
	return params.ModelType == modelType && params.InstanceID == instanceId
}
