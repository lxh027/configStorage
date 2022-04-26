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
	Type     int    `json:"type"`
	Username string `json:"username"`
}

type UserNamespace struct {
	UserID      int `json:"user_id"`
	NamespaceID int `json:"namespace_id"`
	Type        int `json:"type"`
}

type Op struct {
	NamespaceID int    `json:"namespace_id"`
	RaftID      string `json:"raft_id"`
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

type AuthQuery struct {
	NamespaceID int    `json:"namespace_id"`
	UserName    string `json:"username"`
	Type        int    `json:"type"`
}
