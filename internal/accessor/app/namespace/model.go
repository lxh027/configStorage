package namespace

type Namespace struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	ClusterID  int    `json:"cluster_id"`
	PrivateKey string `json:"private_key"`
	Owner      int    `json:"owner"`
}
