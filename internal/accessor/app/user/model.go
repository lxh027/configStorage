package user

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  int    `json:"is_admin"`
}
