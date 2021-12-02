package models

import (
	"errors"
	"monitor/model/enum"
)

// EntryModel entry model. to save entry info for each instance
type EntryModel struct {
	Params
	entryInfo EntryInfo
}

// EntryInfo contains 3 params.
type EntryInfo struct {
	// lastAppend indicates the last entry id received from leader
	LastAppend uint32
	// commitIndex indicates the last entry id labeled as committed
	CommitIndex uint32
	// lastApplied indicates the last entry id been applied to state machine
	LastApplied uint32
}

// NewEntryModel gen new entry model
func NewEntryModel(instanceId uint32) *EntryModel {
	entryModel := EntryModel{
		Params: Params{
			ModelType:  enum.Entry,
			InstanceID: instanceId,
			Title:      "entry_info",
		},
		entryInfo: EntryInfo{0, 0, 0},
	}
	return &entryModel
}

// Update receive data as EntryInfo
func (model *EntryModel) Update(param uint32, data interface{}) error {
	d, err := data.(EntryInfo)
	if !err {
		return errors.New("parse entry info data error")
	}
	model.entryInfo = d
	return nil
}

// Save no save action should be defined
func (model *EntryModel) Save(path string) error {
	return nil
}

func (model *EntryModel) This() interface{} {
	return model
}

func (model *EntryModel) GetParams() *Params {
	return &model.Params
}

func (model *EntryModel) GetData() interface{} {
	return model.entryInfo
}
