package log

type Log struct {
	ID          int    `json:"id"`
	Type        int    `json:"type"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	NamespaceID int    `json:"namespace_id"`
	UserID      int    `json:"user_id"`
	Status      int    `json:"status"`
}
