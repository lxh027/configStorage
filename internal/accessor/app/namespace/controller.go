package namespace

import "C"
import (
	"configStorage/internal/accessor/app/user"
	"configStorage/internal/accessor/global"
	"configStorage/tools/formatter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	namespaceService Service
)

func GetUserNamespaces(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query Query
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	var namespaces []WithAuth
	var err error
	if namespaces, err = namespaceService.GetUserNamespaces(userId, query.Name, query.PageSize*(query.PageIndex-1), query.PageSize); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", map[string]interface{}{
		"count":      len(namespaces),
		"namespaces": namespaces,
	}))
}

func NewNamespace(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query NewNamespaceQuery
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := namespaceService.NewNamespace(userId, query.Name, query.RaftID); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}

func SetAuth(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query AuthQuery
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := namespaceService.AuthUserPrivileges(userId, query.UserName, query.NamespaceID, query.Type); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}

func DeleteNamespace(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query Op
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := namespaceService.DeleteNamespace(userId, query.NamespaceID); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}

func UpdateNamespaceRaftID(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query Op
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := namespaceService.UpdateNamespaceRaftID(userId, query.NamespaceID, query.RaftID); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}
