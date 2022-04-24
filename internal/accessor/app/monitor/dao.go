package monitor

import "configStorage/internal/accessor/global"

type Dao struct{}

func (dao *Dao) GetClusterPeers(raftID string) ([]Peer, error) {
	var peers []Peer
	err := global.MysqlClient.
		Table(
			"(?) as a",
			global.MysqlClient.Model(&Monitor{}).
				Select([]string{"peer_id", "is_leader"}).
				Where("raft_id = ?", raftID).
				Limit(100).Order("id desc")).
		Group("peer_id").Find(&peers).Error
	return peers, err
}

func (dao *Dao) GetPeerData(raftID string, peerID int) ([]Monitor, error) {
	var monitors []Monitor
	err := global.MysqlClient.Where("raft_id = ? AND peer_id = ?", raftID, peerID).
		Order("id desc").Find(&monitors).Error
	return monitors, err
}