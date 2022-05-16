package user

import (
	"configStorage/internal/accessor/global"
	"configStorage/tools/md5"
)

type Service struct {
	userDao Dao
}

func (s *Service) Login(user *User) (int, bool) {
	passwordMd5 := md5.GetMd5(user.Password)
	user.Password = passwordMd5
	return s.userDao.CheckUserPassword(user)
}

func (s *Service) AddUser(user User) bool {
	if s.userDao.CheckUserExisted(user.Username) {
		global.Log.Printf("user already existed\n")
		return false
	}
	user.Password = md5.GetMd5(user.Password)
	if err := s.userDao.AddUser(user); err != nil {
		global.Log.Printf("error occurs when insert user: %v", err.Error())
		return false
	}
	return true
}

func (s *Service) GetUsers(username string, offset, limit int) ([]User, error) {
	return s.userDao.GetUsersByName(username, limit, offset)
}

func (s *Service) SetAdmin(userID int) bool {
	return s.userDao.SetAdmin(userID)
}
