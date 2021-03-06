package user

import (
	"configStorage/internal/accessor/global"
)

type Dao struct{}

func (dao *Dao) AddUser(user User) error {
	return global.MysqlClient.Create(&user).Error
}

func (dao *Dao) UpdatePassword(id int, password string) error {
	return global.MysqlClient.Model(User{}).Where("id = ?", id).Update("password", password).Error
}

func (dao *Dao) DelUser(id int) error {
	return global.MysqlClient.Where("id = ?", id).Delete(&User{}).Error
}

func (dao *Dao) GetUserByID(id int) (User, error) {
	var user User
	err := global.MysqlClient.Where("id = ?", id).First(&user).Error
	return user, err
}

func (dao *Dao) GetAUserByName(username string) (User, error) {
	var user User
	err := global.MysqlClient.Select([]string{"id", "username", "is_admin"}).
		Where(`username = ?`, username).
		First(&user).Error
	return user, err
}

func (dao *Dao) GetUsersByName(username string, limit int, offset int) ([]User, error) {
	var users []User
	err := global.MysqlClient.Select([]string{"id", "username", "is_admin"}).
		Where(`username LIKE "%` + username + `%"`).
		Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (dao *Dao) GetAllUsers(limit int, offset int) ([]User, error) {
	var users []User
	err := global.MysqlClient.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (dao *Dao) CheckUserPassword(user *User) (int, bool) {
	err := global.MysqlClient.Where("username = ? AND password = ?", user.Username, user.Password).First(user).Error
	return user.ID, err == nil
}

func (dao *Dao) CheckUserExisted(username string) bool {
	err := global.MysqlClient.Where("username = ?", username).First(&User{}).Error
	return err == nil
}

func (dao *Dao) CheckUserAdmin(userID int) bool {
	err := global.MysqlClient.Where("use_id = ? AND is_admin = 1", userID).First(&User{}).Error
	return err == nil
}

func (dao *Dao) SetAdmin(userID int) bool {
	err := global.MysqlClient.Model(User{}).Where("id = ?", userID).Update("is_admin", 1).Error
	return err == nil
}
