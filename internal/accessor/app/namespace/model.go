package namespace

type Namespace struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	RaftID     string `json:"raft_id"`
	UserID     int    `json:"user_id"`
	PrivateKey string `json:"private_key,omitempty"`
}

type WithAuth struct {
	Namespace
	Type int `json:"type"`
}

type UserNamespace struct {
	UserID      int `json:"user_id"`
	NamespaceID int `json:"namespace_id"`
	Type        int `json:"type"`
}

type Query struct {
	Name      string `json:"name"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type NewNamespaceQuery struct {
	Name   string `json:"name"`
	RaftID string `json:"raft_id"`
}
