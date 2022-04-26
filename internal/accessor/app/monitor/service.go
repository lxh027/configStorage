package monitor

import (
	"configStorage/internal/accessor/app/cluster"
	"errors"
)

type Service struct {
	clusterDao cluster.Dao
	monitorDao Dao
}

func (s *Service) GetRaftPeers(userID int, raftID string) ([]Peer, error) {
	if !s.clusterDao.CheckUserCluster(userID, raftID) {
		return nil, errors.New("user don't have cluster's authorization")
	}
	return s.monitorDao.GetClusterPeers(raftID)
}

func (s *Service) GetPeerMonitorData(userID int, raftID string, peerID int) (*Data, error) {
	if !s.clusterDao.CheckUserCluster(userID, raftID) {
		return nil, errors.New("user don't have cluster's authorization")
	}
	var monitors []Monitor
	var err error
	if monitors, err = s.monitorDao.GetPeerData(raftID, peerID); err != nil {
		return nil, err
	} else if len(monitors) == 0 {
		return nil, errors.New("no monitor data found")
	}

	data := Data{
		Basic: BasicData{
			RaftID:        monitors[0].RaftID,
			PeerID:        monitors[0].PeerID,
			IsLeader:      monitors[0].IsLeader,
			Status:        monitors[0].Status,
			ConfigVersion: monitors[0].ConfigVersion,
		},
		Memory: MemoryData{
			MemoryTotal:     monitors[0].MemoryTotal,
			MemoryCur:       monitors[0].MemoryCur,
			MemoryUsed:      monitors[0].MemoryUsed,
			MemoryAvailable: monitors[0].MemoryAvailable,
		},
		Time: TimeData{
			RevoteTime: monitors[0].RevoteTime,
			CommitTime: monitors[0].CommitTime,
		},
		Raft: make([]RaftData, 0),
	}

	var curTerm, curIdx, cmtIdx, status RaftData = RaftData{Item: "current_term", Data: make([][]int64, 0)},
		RaftData{Item: "current_index", Data: make([][]int64, 0)},
		RaftData{Item: "commit_index", Data: make([][]int64, 0)},
		RaftData{Item: "raft_status", Data: make([][]int64, 0)}

	for i := len(monitors) - 1; i >= 0; i-- {
		monitor := monitors[i]
		timestamp := monitor.Time.UnixMilli()
		curTerm.Data = append(curTerm.Data, []int64{timestamp, int64(monitor.CurrentTerm)})
		curIdx.Data = append(curIdx.Data, []int64{timestamp, int64(monitor.CurrentIndex)})
		cmtIdx.Data = append(cmtIdx.Data, []int64{timestamp, int64(monitor.CommitIndex)})
		status.Data = append(status.Data, []int64{timestamp, int64(monitor.Status)})
	}
	data.Raft = append(data.Raft, curTerm, curIdx, cmtIdx, status)
	return &data, nil
}
