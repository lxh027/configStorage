package main

import (
	"configStorage/internal/scheduler"
	"configStorage/pkg/config"
	"flag"
	"path"
)

func main() {
	var env string
	var readPersist bool
	flag.StringVar(&env, "env", "dev", "配置环境")
	flag.BoolVar(&readPersist, "persist", true, "读取持久化数据")
	flag.Parse()

	p := path.Join("./config", env, "register.yml")
	rp := path.Join("./config", env, "redis.yml")

	registerConfig := scheduler.NewRegisterConfig(p)
	raftConfig := config.NewRedisConfig(rp)

	registerCenter := scheduler.NewRegisterCenter(registerConfig, raftConfig, readPersist)

	registerCenter.Start()
}
