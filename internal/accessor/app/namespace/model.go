package namespace

type Namespace struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	RaftID     string `json:"raft_id"`
	UserID     int    `json:"user_id"`
	PrivateKey string `json:"private_key,omitempty"`
}

type UserNamespace struct {
	UserID      int `json:"user_id"`
	NamespaceID int `json:"namespace_id"`
	Type        int `json:"type"`
}
