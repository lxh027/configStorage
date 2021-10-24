package grpc_server

import (
	"fmt"
	"google.golang.org/grpc"
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
	/*go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			panic("grpc server start error")
		}
	}()*/
	_ = grpcServer.Serve(lis)
}

