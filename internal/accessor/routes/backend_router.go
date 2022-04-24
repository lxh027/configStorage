package routes

import (
	"configStorage/internal/accessor/app/cluster"
	"configStorage/internal/accessor/app/log"
	"configStorage/internal/accessor/app/monitor"
	"configStorage/internal/accessor/app/namespace"
	"configStorage/internal/accessor/app/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BackendRoutes(router *gin.Engine) {
	usr := router.Group("user")
	{
		usr.POST("register", user.Register)
		usr.POST("login", user.Login)
		usr.POST("logout", user.Logout)
		usr.POST("getUsers", user.GetUsers)
		usr.POST("setAdmin", user.SetAdmin)
	}

	clt := router.Group("cluster")
	{
		clt.POST("getAllClusters", cluster.GetAllCluster)
		clt.POST("getUserClusters", cluster.GetUserCluster)
		clt.POST("authUserCluster", cluster.AuthUserCluster)
		clt.POST("unAuthUserCluster", cluster.UnAuthUserCluster)
	}

	name := router.Group("namespace")
	{
		name.POST("getUserNamespaces", namespace.GetUserNamespaces)
		name.POST("newNamespace", namespace.NewNamespace)
		name.POST("setAuth", namespace.SetAuth)
	}

	lg := router.Group("log")
	{
		lg.POST("getAllLogs", log.GetAllLogs)
		lg.POST("newLog", log.NewLog)
		lg.POST("delLog", log.DelLog)
		lg.POST("commit", log.Commit)
	}

	cfg := router.Group("config")
	{
		cfg.POST("getAllConfigs", log.GetALlConfigs)
	}

	mnt := router.Group("monitor")
	{
		mnt.POST("getRaftPeers", monitor.GetRaftPeers)
		mnt.POST("getPeerData", monitor.GetRaftPeerMonitorData)
	}

	router.StaticFS("/admin", http.Dir("./web"))
}
