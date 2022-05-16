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

func GetUsers(c *gin.Context) {
	var query userQuery
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	var users []User
	var err error
	if users, err = userService.GetUsers(query.Username, query.PageSize*(query.PageIndex-1), query.PageSize); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", map[string]interface{}{
		"count": len(users),
		"users": users,
	}))
}

func SetAdmin(c *gin.Context) {
	var user User
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if !userService.SetAdmin(user.ID) {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "set user admin error", nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}

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
	if v := session.Get(LoginSession); v != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user already logged in", v))
		return
	}

	var user User
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	var userID int
	var ok bool
	if userID, ok = userService.Login(&user); !ok {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user or password error", nil))
		return
	}

	session.Set(LoginSession, userID)
	if user.IsAdmin == Admin {
		session.Set(AdminSession, 1)
	}
	session.Save()

	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "login success", user.IsAdmin))
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "logout success", nil))
}
