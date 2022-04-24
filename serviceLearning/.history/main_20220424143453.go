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

		当使用自签证书时，代码变成这个样子
		服务端代码为：
		creds,err:= credentials.NewServerTLSFromFile("keys/server.crt","keys/Server_no_passwd.key")//加载文件，分别是服务端证书存放的位置，和服务端的私匙存放的位置
		if err!=nil {
				log.Fatal(err)
			}
		rpcServer := grpc.NewServer(grpc.Creds(creds))
		客户端代码为：
		creds,err := credentials.NewClientTLSFromFile("keys/server.crt","jtthink.com")//两个参数，前者是证书的存放位置，后者是服务端的serverNameOverride（再生成证书的时候有这个参数）
			if err!=nil {
				log.Fatal(err)
			}
		conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(creds))
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
