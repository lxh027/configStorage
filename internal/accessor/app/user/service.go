package user

import (
	"configStorage/internal/accessor/global"
	"configStorage/tools/md5"
)

type Service struct {
	userDao Dao
}

func (s *Service) Login(user User) bool {
	passwordMd5 := md5.GetMd5(user.Password)
	return s.userDao.CheckUserPassword(user.Username, passwordMd5)
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
