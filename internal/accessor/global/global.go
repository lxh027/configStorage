package global

import (
	"configStorage/internal/accessor/db"
	"configStorage/internal/scheduler"
	config2 "configStorage/pkg/config"
	"configStorage/pkg/redis"
	"gorm.io/gorm"
	"path"
)

const CodeError = -1
const CodeSuccess = 0

var MysqlClient *gorm.DB
var RedisClient *redis.Client

var SDBClient scheduler.Client

var Log Logger

var DatabaseCfg config2.Database
var LogCfg config2.Log
var RedisCfg config2.Redis
var ServerCfg config2.Server
var SessionCfg config2.Session

func Init(env string) {
	basePath := path.Join("./config", env)

	databasePath := path.Join(basePath, "database.yml")
	DatabaseCfg = config2.NewDatabaseConfig(databasePath)

	logPath := path.Join(basePath, "log.yml")
	LogCfg = config2.NewLogConfig(logPath)

	redisPath := path.Join(basePath, "redis.yml")
	RedisCfg = config2.NewRedisConfig(redisPath)

	serverPath := path.Join(basePath, "server.yml")
	ServerCfg = config2.NewServerConfig(serverPath)

	sessionPath := path.Join(basePath, "session.yml")
	SessionCfg = config2.NewSessionConfig(sessionPath)

	MysqlClient = db.NewMysqlConn(&DatabaseCfg)
	RedisClient, _ = redis.NewRedisClient(&RedisCfg)
	SDBClient = scheduler.NewSchedulerClient(ServerCfg.SchedulerAddr)

	Log = NewLogger()
}
