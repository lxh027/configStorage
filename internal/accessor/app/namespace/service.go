package namespace

import (
	"configStorage/internal/accessor/global"
	"configStorage/tools/random"
	"errors"
)

type Service struct {
	namespaceDao Dao
}

func (s *Service) NewNamespace(userId int, name string, raftID string) error {
	privateKey := random.RandString(16)
	if err := global.SDBClient.NewNamespace(name, privateKey, raftID); err != nil {
		return err
	}
	return s.namespaceDao.NewNamespace(userId, name, raftID, privateKey)
}

// UpdateNamespace TODO updateNamespace
func (s *Service) UpdateNamespace() {}

// DeleteNamespace TODO deleteNamespace
func (s *Service) DeleteNamespace() {}

func (s *Service) AuthUserPrivileges(me int, userID int, namespaceID int, tp int) error {
	if s.namespaceDao.CheckPriv(me, namespaceID) != Owner {
		return errors.New("not namespace's owner")
	}
	return s.namespaceDao.SetUserPrivileges(userID, namespaceID, tp)
}

func (s *Service) GetUserNamespaces(userID int, name string, offset, limit int) ([]WithAuth, error) {
	return s.namespaceDao.GetUserNamespace(userID, name, offset, limit)
}
