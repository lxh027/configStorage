package global

import (
	"configStorage/internal/accessor/config"
	"configStorage/internal/accessor/db"
	"configStorage/internal/accessor/redis"
	"configStorage/internal/scheduler"
	"gorm.io/gorm"
	"path"
)

const CodeError = -1
const CodeSuccess = 0

var MysqlClient *gorm.DB
var RedisClient *redis.Client

var SDBClient scheduler.Client

var Log Logger

var DatabaseCfg config.Database
var LogCfg config.Log
var RedisCfg config.Redis
var ServerCfg config.Server
var SessionCfg config.Session

func Init(env string) {
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

	MysqlClient = db.NewMysqlConn(&DatabaseCfg)
	RedisClient = redis.NewRedisClient(&RedisCfg)

	SDBClient = scheduler.NewSchedulerClient(ServerCfg.SchedulerAddr)

	Log = NewLogger()
}
