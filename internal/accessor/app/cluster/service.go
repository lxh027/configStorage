package cluster

import (
	"configStorage/internal/accessor/app/user"
	"configStorage/internal/accessor/global"
)

type Service struct {
	userDao    user.Dao
	clusterDao Dao
}

func (s *Service) FetchClusters() ([]Cluster, error) {
	cs, err := global.SDBClient.GetClusters()
	if err != nil {
		return nil, err
	}

	clusters := make([]Cluster, 0)
	for _, c := range cs {
		clusters = append(clusters, Cluster{c.RaftID, c.Address})
	}

	if s.clusterDao.AddClusters(clusters) != nil {
		return clusters, err
	}

	return clusters, nil
}

func (s *Service) GetUserClusters(userID int) ([]Cluster, error) {
	return s.clusterDao.GetUserClusters(userID)
}

func (s *Service) AuthorizeClusters(userID int, raftId string) error {
	return s.clusterDao.AddUserCluster(userID, raftId)
}

func (s *Service) UnAuthorizeClusters(userID int, raftId string) error {
	return s.clusterDao.DeleteUserCluster(userID, raftId)
}

func (s *Service) CheckUserCluster(userID int, raftID string) bool {
	return s.clusterDao.CheckUserCluster(userID, raftID)
}
