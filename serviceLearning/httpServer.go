package main

import (
	"context"
	"log"
	"net/http"
	"serviceLearning/helper"
	"serviceLearning/services"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx) //下面这两句不是必须的
	defer cancel()

	gwmux := runtime.NewServeMux() //路由

	// opts := []grpc.DialOption{grpc.WithInsecure()}//不使用证书
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}//不使用证书

	opts := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := services.RegisterProdServiceHandlerFromEndpoint(ctx,
		gwmux, "localhost:8081", opts) //endpoint指的是grpc对应的地址
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux, //使用已给的路由
	}
	log.Printf("server listening at %v", httpServer.Addr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	log.Printf("server over!!")
}
