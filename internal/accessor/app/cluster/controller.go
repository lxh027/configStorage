package cluster

import (
	"configStorage/internal/accessor/app/user"
	"configStorage/internal/accessor/global"
	"configStorage/tools/formatter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	clusterService Service
)

func GetAllCluster(c *gin.Context) {
	session := sessions.Default(c)
	if v := session.Get(user.AdminSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not admin", v))
		return
	}

	var clusters []Cluster
	var err error
	if clusters, err = clusterService.FetchClusters(); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", map[string]interface{}{
		"count":   len(clusters),
		"cluster": clusters,
	}))
}

func GetUserCluster(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query UserCluster
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if query.UserID != 0 {
		userId = query.UserID
	}
	
	var clusters []Cluster
	var err error
	if clusters, err = clusterService.GetUserClusters(userId); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", map[string]interface{}{
		"count":   len(clusters),
		"cluster": clusters,
	}))
}

func AuthUserCluster(c *gin.Context) {
	session := sessions.Default(c)
	if v := session.Get(user.AdminSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not admin", v))
		return
	}

	var query UserCluster
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := clusterService.AuthorizeClusters(query.UserID, query.RaftID); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}

func UnAuthUserCluster(c *gin.Context) {
	session := sessions.Default(c)
	if v := session.Get(user.AdminSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not admin", v))
		return
	}

	var query UserCluster
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := clusterService.UnAuthorizeClusters(query.UserID, query.RaftID); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}
