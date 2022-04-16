package log

import (
	"configStorage/internal/accessor/app/namespace"
	"configStorage/internal/accessor/global"
	"configStorage/internal/scheduler"
	"errors"
)

type Service struct {
	namespaceDao namespace.Dao
	logDao       Dao
}

func (s *Service) GetLogs(userID int, namespaceID int, offset, limit int) ([]Log, error) {
	priv := s.namespaceDao.CheckPriv(userID, namespaceID)
	if priv == namespace.Abandon {
		return nil, errors.New("not authorized")
	}
	return s.logDao.GetLogs(namespaceID, offset, limit)

}

func (s *Service) AddLog(userID int, namespaceID int, key string, value string, tp int) error {
	priv := s.namespaceDao.CheckPriv(userID, namespaceID)
	if priv != namespace.Owner && priv != namespace.Normal {
		return errors.New("not authorized")
	}
	return s.logDao.AddLog(userID, namespaceID, key, value, tp)
}

func (s *Service) DelLog(userID int, logID int) error {
	if log, err := s.logDao.GetLogByID(logID); err != nil {
		return err
	} else if log.Status == Committed {
		return errors.New("log is committed")
	} else {
		priv := s.namespaceDao.CheckPriv(userID, log.NamespaceID)
		if priv != namespace.Owner && priv != namespace.Normal {
			return errors.New("not authorized")
		}
	}
	return s.logDao.DelLog(logID)
}

func (s *Service) Commit(userID int, namespaceId int, logId int, lastCommitId int) error {
	priv := s.namespaceDao.CheckPriv(userID, namespaceId)
	if priv != namespace.Owner && priv != namespace.Normal {
		return errors.New("not authorized")
	}

	np := s.namespaceDao.GetNamespace(namespaceId)
	if np == nil {
		return errors.New("namespace not existed")
	}
	// TODO lock

	l, err := s.logDao.GetLogsForRange(namespaceId, lastCommitId+1, logId)
	if err != nil {
		return err
	}
	logs := make([]scheduler.Log, 0)
	for _, item := range l {
		logs = append(logs, scheduler.Log{
			ID:    item.ID,
			Type:  item.Type,
			Key:   item.Key,
			Value: item.Value,
		})
	}
	id, err := global.SDBClient.Commit(np.Name, np.PrivateKey, logs)
	if err != nil {
		global.Log.Printf("commit log error for namespace %v: %v", namespaceId, err.Error())
	}
	return s.logDao.Commit(namespaceId, lastCommitId+1, id)
}

func (s *Service) Restore(userID int, namespaceId int, logID int, lastCommitId int) error {
	priv := s.namespaceDao.CheckPriv(userID, namespaceId)
	if priv != namespace.Owner && priv != namespace.Normal {
		return errors.New("not authorized")
	}
	np := s.namespaceDao.GetNamespace(namespaceId)
	if np == nil {
		return errors.New("namespace not existed")
	}
	// TODO lock

	l, err := s.logDao.GetLogsForRange(namespaceId, logID+1, lastCommitId)
	if err != nil {
		return err
	}
	logs := make([]scheduler.Log, 0)
	for i := len(l); i >= 0; i-- {
		logs = append(logs, scheduler.Log{
			ID:    l[i].ID,
			Type:  1 - l[i].Type,
			Key:   l[i].Key,
			Value: l[i].Value,
		})
	}
	id, err := global.SDBClient.Commit(np.Name, np.PrivateKey, logs)
	if err != nil {
		global.Log.Printf("restore log error for namespace %v: %v", namespaceId, err.Error())
	}
	return s.logDao.Commit(namespaceId, id+1, lastCommitId)
}

func (s *Service) LastCommittedID(namespaceID int) int {
	return s.logDao.LastCommittedID(namespaceID)
}
