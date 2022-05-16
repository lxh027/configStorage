package monitor

import (
	"configStorage/internal/accessor/app/user"
	"configStorage/internal/accessor/global"
	"configStorage/tools/formatter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	monitorService Service
)

func GetRaftPeers(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query Peer
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}
	isAdmin := session.Get(user.AdminSession) != nil

	var peers []Peer
	var err error
	if peers, err = monitorService.GetRaftPeers(userId, query.RaftID, isAdmin); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", peers))

}

func GetRaftPeerMonitorData(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	isAdmin := session.Get(user.AdminSession) != nil

	var query Peer
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	var data *Data
	var err error
	if data, err = monitorService.GetPeerMonitorData(userId, query.RaftID, query.PeerID, isAdmin); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", data))
}
