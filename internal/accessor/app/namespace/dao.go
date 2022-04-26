package namespace

import (
	"configStorage/internal/accessor/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Dao struct{}

func (dao *Dao) NewNamespace(userId int, name string, raftID string, privateKey string) error {
	n := Namespace{
		Name:       name,
		PrivateKey: privateKey,
		UserID:     userId,
		RaftID:     raftID,
	}
	return global.MysqlClient.Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&n).Error; e != nil {
			return e
		}
		return tx.Create(UserNamespace{NamespaceID: n.ID, UserID: userId, Type: Owner}).Error
	})
}

func (dao *Dao) GetUserNamespace(userId int, name string, offset, limit int) ([]WithAuth, error) {
	var namespaces []WithAuth
	err := global.MysqlClient.Model(Namespace{}).
		Select("namespace.id, namespace.name, namespace.raft_id, namespace.user_id, namespace.private_key, user_namespace.type, user.username").
		Joins("join user_namespace on namespace.id = user_namespace.namespace_id").
		Joins("join user on user.id = user_namespace.user_id").
		Where("user_namespace.user_id = ? AND namespace.name LIKE '%"+name+"%'", userId).
		Limit(limit).Offset(offset).
		Find(&namespaces).
		Error
	return namespaces, err
}

func (dao *Dao) SetUserPrivileges(userID int, namespaceID int, priv int) error {
	return global.MysqlClient.Clauses(clause.OnConflict{UpdateAll: true}).
		Create(UserNamespace{UserID: userID, NamespaceID: namespaceID, Type: priv}).
		Error
}

func (dao *Dao) CheckPriv(userId int, namespaceID int) int {
	var un UserNamespace
	err := global.MysqlClient.Where("user_id = ? AND namespace_id = ?", userId, namespaceID).
		First(&un).Error
	if err != nil {
		return Abandon
	}
	return un.Type
}

func (dao *Dao) GetNamespace(namespaceID int) *Namespace {
	var np Namespace
	if global.MysqlClient.Where("id = ?", namespaceID).First(&np).Error != nil {
		return nil
	}
	return &np
}

func (dao *Dao) GetNamespaceWithAuth(namespaceID int, userID int) (WithAuth, error) {
	var namespace WithAuth
	err := global.MysqlClient.Model(Namespace{}).
		Select("namespace.id, namespace.name, namespace.raft_id, namespace.user_id, namespace.private_key, user_namespace.type, user.username").
		Joins("join user_namespace on namespace.id = user_namespace.namespace_id").
		Joins("join user on user.id = user_namespace.user_id").
		Where("user_namespace.user_id = ? AND namespace.id = ?", userID, namespaceID).
		First(&namespace).
		Error
	return namespace, err
}

func (dao *Dao) UpdateRaftID(namespaceID int, raftID string) error {
	return global.MysqlClient.Model(&Namespace{}).Where("id = ?", namespaceID).Update("raft_id", raftID).Error
}

func (dao *Dao) DeleteNamespace(namespaceID int) error {
	return global.MysqlClient.Where("id = ?", namespaceID).Delete(&Namespace{}).Error
}
