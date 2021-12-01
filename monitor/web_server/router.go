package web_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// registerRoutes to register controllers to routes
func registerRoutes(router *gin.Engine)  {

	router.StaticFS("/", http.Dir("./web/"))
}
