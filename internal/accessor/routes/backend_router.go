package routes

import (
	"configStorage/internal/accessor/app/cluster"
	"configStorage/internal/accessor/app/namespace"
	"configStorage/internal/accessor/app/user"
	"github.com/gin-gonic/gin"
)

func BackendRoutes(router *gin.Engine) {
	usr := router.Group("user")
	{
		usr.POST("register", user.Register)
		usr.POST("login", user.Login)
		usr.POST("logout", user.Logout)
		usr.POST("getUsers", user.GetUsers)
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
	}

	//router.StaticFS("/admin", http.Dir("./web"))
}
