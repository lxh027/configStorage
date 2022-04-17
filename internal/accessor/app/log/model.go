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

type FetchQuery struct {
	NamespaceID int `json:"namespace_id"`
	PageIndex   int `json:"pageIndex"`
	PageSize    int `json:"pageSize"`
}

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
