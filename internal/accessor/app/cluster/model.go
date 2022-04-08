package cluster

type Cluster struct {
	RaftID  string `json:"raft_id"`
	Address string `json:"address"`
}

type UserCluster struct {
	UserID int    `json:"user_id"`
	RaftID string `json:"raft_id"`
}
