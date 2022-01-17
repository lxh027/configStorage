package user

import (
	"configStorage/internal/accessor/global"
	"configStorage/tools/formatter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	userService Service
)

func Register(c *gin.Context) {
	var user User
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if !userService.AddUser(user) {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "add user error", nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "add user success", nil))
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	if v := session.Get("username"); v != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user already logged in", v))
		return
	}

	var user User
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if !userService.Login(user) {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user or password error", nil))
		return
	}

	session.Set("username", user.Username)
	session.Save()

	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "login success", nil))
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "logout success", nil))
}
