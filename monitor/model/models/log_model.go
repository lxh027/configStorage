package models

import (
	"errors"
	"log"
	"monitor/model/enum"
	"os"
)

// LogModel log model. to receive and save logs from instances
// save to file periodically
type LogModel struct {
	Params
	logs	LogModelData
}

type LogModelData []string

// NewLogModel gen log model
func NewLogModel(instanceId uint32) *LogModel {
	logModel := LogModel{
		Params: Params{
			ModelType: enum.Log,
			InstanceID: instanceId,
			Title: "log",
		},
		logs: make(LogModelData, 0),
	}
	return &logModel
}

// Update receive data as LogModelData
func (model *LogModel) Update(param uint8, data interface{}) error {
	d, err := data.(LogModelData)
	if !err {
		return errors.New("parse log model data error")
	}
	model.logs = append(model.logs, d...)
	return nil
}

func (model *LogModel) Save(path string) error {
	// open log file with append && create mode
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Printf("open or create file error, file path: %v", err.Error())
		return errors.New("open or create file error")
	}
	for _, l := range model.logs {
		_, _ = file.WriteString(l)
	}
	return nil
}

func (model *LogModel) This() interface{} {
	return model
}

func (model *LogModel) GetParams() *Params  {
	return &model.Params
}

func (model *LogModel) GetData() interface{} {
	return model.logs
}

