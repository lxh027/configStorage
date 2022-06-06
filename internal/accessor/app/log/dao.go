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

func (dao *Dao) DelLog(logID int) error {
	return global.MysqlClient.Where("id = ?", logID).Delete(&Log{}).Error
}

func (dao *Dao) GetLogs(namespaceID, offset, limit int) ([]Log, error) {
	var logs []Log
	err := global.MysqlClient.Where("namespace_id = ?", namespaceID).Order("id desc").
		Offset(offset).Limit(limit).Find(&logs).Error
	return logs, err
}

func (dao *Dao) GetLogsWithUsername(namespaceID, offset, limit int) ([]WithUsername, error) {
	var logs []WithUsername

	err := global.MysqlClient.Model(Log{}).
		Select("log.id, log.type, log.key, log.value, log.namespace_id, log.status, user.username").
		Joins("join user on user.id = log.user_id").
		Where("log.namespace_id = ?", namespaceID).Order("log.id desc").
		Offset(offset).Limit(limit).Find(&logs).Error
	return logs, err
}

func (dao *Dao) GetLogByID(logID int) (Log, error) {
	var log Log
	err := global.MysqlClient.Where("id = ?", logID).First(&log).Error
	return log, err
}

func (dao *Dao) GetLogsForRange(namespaceID int, lower, upper int) ([]Log, error) {
	var logs []Log
	err := global.MysqlClient.
		Where("namespace_id = ? AND id >= ? AND id <= ?", namespaceID, lower, upper).
		Find(&logs).Error
	return logs, err
}

func (dao *Dao) Commit(namespaceID, lower, upper int) error {
	return global.MysqlClient.Model(&Log{}).
		Where("namespace_id = ? AND id >= ? AND id <= ? AND status = ?", namespaceID, lower, upper, Uncommitted).
		Update("status", Committed).Error
}

func (dao *Dao) Restore(namespaceID, lower, upper int) error {
	return global.MysqlClient.Model(&Log{}).
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
