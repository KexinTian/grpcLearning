package main

import (
	"log"
	"net"
	"serviceLearning/services"

	"google.golang.org/grpc"
)

func main() {
	/*
		用grpc创建一个grpcServer
	*/
	rpcServer := grpc.NewServer()
	/*
		将自己写的service 注册到 grpcServer 上
		关注注册方法 func RegisterProdServiceServer
		services.ProdService 是自己实现的struct
	*/
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	/*
		创建监听端口
	*/

	lis, _ := net.Listen("tcp", ":8082")
	/*
		让 grpcServer去监听端口
	*/
	log.Printf("server listening at %v", lis.Addr())
	rpcServer.Serve(lis)

}
