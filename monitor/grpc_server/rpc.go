package grpc_server

import (
	"context"
	"monitor/api"
	"monitor/filter"
	"monitor/model/models"
)

func (server *monitorReportServer) InstanceInfo(ctx context.Context, msg *api.InstanceMsg) (*api.EmptyMsg, error)  {
	// append logs
	filter.AppendLogs(msg.InstanceID, msg.Logs)
	// update entry state
	filter.UpdateEntryState(msg.InstanceID, models.EntryInfo{
		LastAppend: msg.RaftStatus.LastReceived,
		CommitIndex: msg.RaftStatus.CommitIndex,
		LastApplied: msg.RaftStatus.LastApplied,
	})
	return &api.EmptyMsg{}, nil
}

func (server *monitorReportServer) LeaderInfo(ctx context.Context, msg *api.LeaderMsg) (*api.EmptyMsg, error) {
	return &api.EmptyMsg{}, nil
}
