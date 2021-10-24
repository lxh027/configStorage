package web_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"monitor/config"
)

// StartServer to start web server
func StartServer()  {
	serverConfig := config.GetServerConfig()
	// set run mode
	gin.SetMode(serverConfig.Web.Mode)
	// init a gin server
	httpServer := gin.Default()
	httpServer.Use(gin.Recovery())
	// register routes
	registerRoutes(httpServer)
	log.Printf("starting web server on %s:%d in %s mode\n", serverConfig.Web.Host, serverConfig.Web.Port, serverConfig.Web.Mode)
	err := httpServer.Run(fmt.Sprintf("%s:%d", serverConfig.Web.Host, serverConfig.Web.Port))
	if err != nil {
		panic("start http server error")
	}
}