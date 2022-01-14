package main

import (
	"configStorage/internal/scheduler"
	"flag"
	"path"
)

func main() {
	var env string

	flag.StringVar(&env, "env", "dev", "配置环境")
	flag.Parse()

	p := path.Join("./config", env, "register.yml")

	registerConfig := scheduler.NewRegisterConfig(p)

	registerCenter := scheduler.NewRegisterCenter(registerConfig)

	registerCenter.Start()
}
