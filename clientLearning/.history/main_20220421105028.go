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
		客户端创建连接的时候默认必须使用加密传输，否则会直接报错
		conn, err := grpc.Dial(":8082")
		在这个简单的例子中，并没有使用加密传输
		grpc.WithInsecure() 表示禁用加密传输
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
