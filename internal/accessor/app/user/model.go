package user

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	IsAdmin  int    `json:"is_admin"`
}

type userQuery struct {
	Username  string `json:"username"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}
