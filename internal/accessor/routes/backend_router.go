package routes

import (
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

	//router.StaticFS("/admin", http.Dir("./web"))
}
