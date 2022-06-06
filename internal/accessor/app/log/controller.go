package log

import (
	"configStorage/internal/accessor/app/user"
	"configStorage/internal/accessor/global"
	"configStorage/tools/formatter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	logService Service
)

func GetAllLogs(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query FetchQuery
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	var logs []WithUsername
	var err error
	if logs, err = logService.GetLogs(userId, query.NamespaceID, query.PageSize*(query.PageIndex-1), query.PageSize); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", map[string]interface{}{
		"count": len(logs),
		"log":   logs,
	}))

}

func GetALlConfigs(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var query FetchQuery
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	var configs []KV
	var err error
	if configs, err = logService.GetConfigs(userId, query.NamespaceID); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", configs))
}

func NewLog(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var log Log
	if c.ShouldBind(&log) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := logService.AddLog(userId, log.NamespaceID, log.Key, log.Value, log.Type); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}

func Commit(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}
	var log Log
	if c.ShouldBind(&log) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	lastCommitID := logService.LastCommittedID(log.NamespaceID)

	if lastCommitID < log.ID {
		// commit
		if err := logService.Commit(userId, log.NamespaceID, log.ID, lastCommitID); err != nil {
			c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
			return
		}
	} else {
		// restore
		if err := logService.Restore(userId, log.NamespaceID, log.ID, lastCommitID); err != nil {
			c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
			return
		}
	}

	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}

func DelLog(c *gin.Context) {
	session := sessions.Default(c)
	var userId int
	if v := session.Get(user.LoginSession); v == nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "user not logged in", v))
		return
	} else {
		userId = v.(int)
	}

	var log Log
	if c.ShouldBind(&log) != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, "param bind error", nil))
		return
	}

	if err := logService.DelLog(userId, log.ID); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, formatter.ApiReturn(global.CodeSuccess, "ok", nil))
}
