package cluster

import (
	"configStorage/internal/accessor/app/user"
	"configStorage/internal/accessor/global"
	"errors"
)

type Service struct {
	userDao    user.Dao
	clusterDao Dao
}

func (s *Service) FetchClusters(me int) ([]Cluster, error) {
	if !s.userDao.CheckUserAdmin(me) {
		return nil, errors.New("only admin permitted")
	}

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

func (s *Service) AuthorizeClusters(me int, userID int, raftId string) error {
	if !s.userDao.CheckUserAdmin(me) {
		return errors.New("only admin permitted")
	}
	return s.clusterDao.AddUserCluster(userID, raftId)
}

func (s *Service) UnAuthorizeClusters(me int, userID int, raftId string) error {
	if !s.userDao.CheckUserAdmin(me) {
		return errors.New("only admin permitted")
	}
	return s.clusterDao.DeleteUserCluster(userID, raftId)
}
