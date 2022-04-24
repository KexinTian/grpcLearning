package main

import (
	"clientLearning/services"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	/*
		客户端创建连接的时候默认使用加密传输，否则会直接报错
		conn, err := grpc.Dial(":8082")
		在这个简单的例子中，并没有使用加密传输,grpc.WithInsecure() 表示禁用加密传输

		当使用自签证书时，代码变成这个样子
		服务端代码为：
		creds,err:= credentials.NewServerTLSFromFile("keys/server.crt","keys/Server_no_passwd.key")//分别是服务端证书存放的位置，和服务端的私匙存放的位置
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

	conn, err := grpc.Dial(":8082", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	/*
		关注 _grpc.pb.go 文件中的ProdServiceClient接口 和 NewProdServiceClient(conn)方法
		其中NewProdServiceClient(conn) 方法 返回的是 ProdServiceClient接口的实例
	*/

	proClient := services.NewProdServiceClient(conn)
	proRes, err := proClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(proRes.ProdStock)
}
