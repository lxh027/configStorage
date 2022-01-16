package accessor

import (
	"configStorage/internal/accessor/config"
	"configStorage/internal/accessor/db"
	"configStorage/internal/accessor/redis"
	"configStorage/internal/accessor/routes"
	"configStorage/tools/formatter"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"os"
	"path"
	"time"
)

var DatabaseCfg config.Database
var LogCfg config.Log
var RedisCfg config.Redis
var ServerCfg config.Server
var SessionCfg config.Session

var MysqlClient *gorm.DB
var RedisClient *redis.Client

func Init(env string) {
	loadCfg(env)
	initMysql(&DatabaseCfg)
	initRedis(&RedisCfg)
}

func loadCfg(env string) {
	basePath := path.Join("./config", env)

	databasePath := path.Join(basePath, "database.yml")
	DatabaseCfg = config.NewDatabaseConfig(databasePath)

	logPath := path.Join(basePath, "log.yml")
	LogCfg = config.NewLogConfig(logPath)

	redisPath := path.Join(basePath, "redis.yml")
	RedisCfg = config.NewRedisConfig(redisPath)

	serverPath := path.Join(basePath, "server.yml")
	ServerCfg = config.NewServerConfig(serverPath)

	sessionPath := path.Join(basePath, "session.yml")
	SessionCfg = config.NewSessionConfig(sessionPath)
}

func initMysql(cfg *config.Database) {
	MysqlClient = db.NewMysqlConn(cfg)
}

func initRedis(cfg *config.Redis) {
	RedisClient = redis.NewRedisClient(cfg)
}

func Run() {
	serverConfig := ServerCfg
	sessionConfig := SessionCfg
	// 运行模式
	gin.SetMode(serverConfig.Mode)
	httpServer := gin.Default()

	// 创建session存储引擎
	sessionStore := cookie.NewStore([]byte(sessionConfig.Key))
	sessionStore.Options(sessions.Options{
		MaxAge: sessionConfig.Age,
		Path:   sessionConfig.Path,
	})
	//设置session中间件
	httpServer.Use(sessions.Sessions(sessionConfig.Name, sessionStore))

	gin.DisableConsoleColor()
	// 生成日志
	logPath := fmt.Sprintf("%s-%s.log", LogCfg.Path, time.Now().String())
	logFile, _ := os.Create(logPath)
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	// 设置日志格式
	httpServer.Use(gin.LoggerWithFormatter(formatter.GetLogFormat))
	httpServer.Use(gin.Recovery())

	// 注册路由
	routes.BackendRoutes(httpServer)

	serverError := httpServer.Run(serverConfig.Host + ":" + serverConfig.Port)

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}
}
