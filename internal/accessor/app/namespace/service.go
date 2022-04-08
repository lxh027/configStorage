package namespace

import (
	"configStorage/tools/random"
	"errors"
)

type Service struct {
	namespaceDao Dao
}

func (s *Service) NewNamespace(userId int, name string, raftID string) error {
	return s.namespaceDao.NewNamespace(userId, name, raftID, random.RandString(16))
}

// UpdateNamespace TODO updateNamespace
func (s *Service) UpdateNamespace() {}

// DeleteNamespace TODO DeleteNamespace
func (s *Service) DeleteNamespace() {}

func (s *Service) AuthUserPrivileges(me int, userID int, namespaceID int, tp int) error {
	if s.namespaceDao.CheckPriv(me, namespaceID) != Owner {
		return errors.New("not namespace's owner")
	}
	return s.namespaceDao.SetUserPrivileges(userID, namespaceID, tp)
}

func (s *Service) GetUserNamespaces(userID int) ([]Namespace, error) {
	return s.namespaceDao.GetUserNamespace(userID)
}
