package log

import "configStorage/internal/accessor/global"

type Dao struct{}

func (dao *Dao) AddLog(userID, namespaceID int, key, value string, tp int) error {
	log := Log{
		UserID:      userID,
		NamespaceID: namespaceID,
		Key:         key,
		Value:       value,
		Type:        tp,
		Status:      Uncommitted,
	}
	return global.MysqlClient.Create(&log).Error
}

func (dao *Dao) GetLogs(namespaceID, offset, limit int) ([]Log, error) {
	var logs []Log
	err := global.MysqlClient.Where("namespace_id = ?", namespaceID).
		Offset(offset).Limit(limit).Find(&logs).Error
	return logs, err
}

func (dao *Dao) GetLogsForRange(namespaceID int, lower, upper int) []Log {
	var logs []Log
	global.MysqlClient.
		Where("namespace_id = ? AND id >= ? AND id <= ?", namespaceID, lower, upper).
		Find(&logs)
	return logs
}

func (dao *Dao) Commit(namespaceID, lower, upper int) error {
	return global.MysqlClient.
		Where("namespace_id = ? AND id >= ? AND id <= ? AND status = ?", namespaceID, lower, upper, Uncommitted).
		Update("status", Committed).Error
}

func (dao *Dao) Restore(namespaceID, lower, upper int) error {
	return global.MysqlClient.
		Where("namespace_id = ? AND id >= ? AND id <= ? AND status = ?", namespaceID, lower, upper, Committed).
		Update("status", Uncommitted).Error
}

func (dao *Dao) LastCommittedID(namespaceID int) int {
	var log Log
	err := global.MysqlClient.Where("namespace_id = ? AND status = ?", namespaceID, Committed).Last(&log).Error
	if err != nil {
		return 0
	}
	return log.ID
}

func (dao *Dao) LastUncommittedID(namespaceID int) int {
	var log Log
	err := global.MysqlClient.Where("namespace_id = ? AND status = ?", namespaceID, Uncommitted).Last(&log).Error
	if err != nil {
		return 0
	}
	return log.ID
}
