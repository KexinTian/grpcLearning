package main

import (
	"log"
	"net"
	"serviceLearning/helper"
	"serviceLearning/services"

	"google.golang.org/grpc"
)

func main_1() {
	creds := helper.GetServerCreds() //双向认证
	/*
		用grpc创建一个grpcServer，不使用证书
		rpcServer := grpc.NewServer()

		当只使用服务器自签证书时，代码变成这个样子
		服务端代码为：
		creds,err:= credentials.NewServerTLSFromFile("keys/server.crt","keys/Server_no_passwd.key")//加载文件，分别是服务端证书存放的位置，和服务端的私钥存放的位置
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
	rpcServer := grpc.NewServer(grpc.Creds(creds))

	/*
		将自己写的service 注册到 grpcServer 上
		关注注册方法 func RegisterProdServiceServer
		services.ProdService 是自己实现的struct
	*/
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	/*
		使用tcp传输
	*/
	lis, _ := net.Listen("tcp", ":8081") //使用tcp协议传输，创建监听端口
	log.Printf("server listening at %v", lis.Addr())
	rpcServer.Serve(lis) //让 grpcServer去监听TCP端口。server会根据传入的lis不同，选择不同的监听模式，最常用的就是 TCPConn，基于 TCP Listener 去做

	/*
		使用http传输
	*/
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.Proto)
	// 	fmt.Println(r.Header)
	// 	fmt.Println(r)
	// 	rpcServer.ServeHTTP(w, r)
	// })
	// httpServer := &http.Server{
	// 	Addr:    ":8081",
	// 	Handler: mux,
	// }
	// log.Printf("server listening at %v", httpServer.Addr)
	// httpServer.ListenAndServeTLS("cert/server.pem", "cert/server.key")

	/*
		在网页上输入“http://localhost:8081/”，显示“gRPC requires HTTP/2”
		在网页上输入“https://localhost:8081/”，显示“此网站无法提供安全连接”,原因是，本服务没有开启TLS验证

		如果只使用服务器端证书的话，这里应该是
		httpServer.ListenAndServeTLS("keys/server.crt","keys/Server_no_passwd.key")
		使用TLS验证后，因为使用的是自签证书，仍然会出现一些错误警告，不过不会影响程序的学习
	*/
}
