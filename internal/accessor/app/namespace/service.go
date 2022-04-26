package namespace

import (
	"configStorage/internal/accessor/app/user"
	"configStorage/internal/accessor/global"
	"configStorage/internal/scheduler"
	"configStorage/tools/random"
	"errors"
)

type Service struct {
	namespaceDao Dao
	userDao      user.Dao
}

func (s *Service) NewNamespace(userId int, name string, raftID string) error {
	privateKey := random.RandString(16)
	if err := global.SDBClient.NewNamespace(name, privateKey, raftID); err != nil {
		return err
	}
	return s.namespaceDao.NewNamespace(userId, name, raftID, privateKey)
}

func (s *Service) UpdateNamespaceRaftID(userID, namespaceID int, raftID string) error {
	if s.namespaceDao.CheckPriv(userID, namespaceID) != Owner {
		return errors.New("not namespace's owner")
	}
	namespace := s.namespaceDao.GetNamespace(namespaceID)
	if namespace == nil {
		return errors.New("namespace not found")
	}
	if err := global.SDBClient.TransNamespace(namespace.Name, namespace.PrivateKey, raftID); err != nil {
		return err
	}
	return s.namespaceDao.UpdateRaftID(namespaceID, raftID)
}

func (s *Service) DeleteNamespace(userId, namespaceId int) error {
	if s.namespaceDao.CheckPriv(userId, namespaceId) != Owner {
		return errors.New("not namespace's owner")
	}
	namespace := s.namespaceDao.GetNamespace(namespaceId)
	if namespace == nil {
		return errors.New("namespace not found")
	}
	if err := global.SDBClient.DeleteNamespace(namespace.Name, namespace.PrivateKey); err != nil && err != scheduler.NamespaceNotExistedErr {
		return err
	}
	return s.namespaceDao.DeleteNamespace(namespaceId)
}

func (s *Service) AuthUserPrivileges(me int, username string, namespaceID int, tp int) error {
	if s.namespaceDao.CheckPriv(me, namespaceID) != Owner {
		return errors.New("not namespace's owner")
	}
	u, err := s.userDao.GetAUserByName(username)
	if err != nil {
		return err
	}
	if me == u.ID {
		return errors.New("can't auth self")
	}
	return s.namespaceDao.SetUserPrivileges(u.ID, namespaceID, tp)
}

func (s *Service) GetUserNamespaces(userID int, name string, offset, limit int) ([]WithAuth, error) {
	return s.namespaceDao.GetUserNamespace(userID, name, offset, limit)
}
