package main

import (
	"monitor/grpc_server"
	"monitor/model"
)

func main() {
	model.StartStorage()
	grpc_server.StartServer()
}