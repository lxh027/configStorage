package main

import (
	"monitor/grpc_server"
	"monitor/model"
	"monitor/web_server"
)

func main() {
	model.StartStorage()
	grpc_server.StartServer()
	web_server.StartServer()
}
