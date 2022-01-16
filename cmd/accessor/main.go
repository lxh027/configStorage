package main

import (
	"configStorage/internal/accessor"
	"flag"
)

func main() {
	defer func() {
		_ = accessor.RedisClient.Client.Close()
		db, _ := accessor.MysqlClient.DB()
		_ = db.Close()
	}()

	var env string

	flag.StringVar(&env, "env", "dev", "配置环境")
	flag.Parse()

	accessor.Init(env)
	accessor.Run()
}
