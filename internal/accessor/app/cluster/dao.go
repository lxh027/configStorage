package cluster

import (
	"configStorage/internal/accessor/global"
	"gorm.io/gorm/clause"
)

type Dao struct{}

func (dao *Dao) GetUserClusters(userID int) ([]Cluster, error) {
	var clusters []Cluster
	err := global.MysqlClient.Model(Cluster{}).
		Select("cluster.raft_id, cluster.address").
		Joins("join user_cluster on cluster.raft_id = user_cluster.raft_id").
		Where("user_cluster.user_id = ?", userID).
		Find(&clusters).
		Error
	return clusters, err
}

func (dao *Dao) CheckUserCluster(userID int, raftID string) bool {
	err := global.MysqlClient.
		Where("user_id = ? AND raft_id = ?", userID, raftID).
		First(&UserCluster{}).Error
	return err == nil
}

func (dao *Dao) AddClusters(clusters []Cluster) error {
	return global.MysqlClient.
		Clauses(clause.OnConflict{DoNothing: true}).
		CreateInBatches(clusters, len(clusters)).
		Error
}

func (dao *Dao) AddUserCluster(userID int, raftId string) error {
	uc := UserCluster{
		UserID: userID,
		RaftID: raftId,
	}
	return global.MysqlClient.Create(&uc).Error
}

func (dao *Dao) DeleteUserCluster(userId int, raftId string) error {
	return global.MysqlClient.Where("user_id = ? AND raft_id = ?", userId, raftId).Delete(&UserCluster{}).Error
}
