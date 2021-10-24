package filter

import (
	"monitor/model"
	"monitor/model/enum"
	"monitor/model/models"
)

// AppendLogs append logs to model
func AppendLogs(instanceID uint32, logs models.LogModelData) {
	if _, err := model.GetSingleModel(enum.Log, instanceID); err != nil {
		model.AddModel(enum.Log, instanceID)
	}
	_ = model.UpdateModel(enum.Log, instanceID, instanceID, logs)
}

// UpdateEntryState update instance's entry state
func UpdateEntryState(instanceID uint32, entryInfo models.EntryInfo) {
	if _, err := model.GetSingleModel(enum.Entry, instanceID); err != nil {
		model.AddModel(enum.Entry, instanceID)
	}
	_ = model.UpdateModel(enum.Entry, instanceID, instanceID, entryInfo)
}
