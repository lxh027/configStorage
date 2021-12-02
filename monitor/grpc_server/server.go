package grpc_server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"monitor/api"
	"monitor/config"
	"net"
)

type monitorReportServer struct {
	api.UnimplementedMonitorReportServer
}

func StartServer() {
	serverConfig := config.GetServerConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverConfig.Grpc.Host, serverConfig.Grpc.Port))
	if err != nil {
		panic("grpc server open tcp port error")
	}
	// opts
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterMonitorReportServer(grpcServer, &monitorReportServer{})
	log.Printf("starting grpc server on %s:%d\n", serverConfig.Grpc.Host, serverConfig.Grpc.Port)
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			panic("grpc server start error")
		}
	}()
	log.Println("grpc server has been started")
}
