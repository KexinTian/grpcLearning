package helper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
)

//获取Server端证书
func GetServerCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key") //读服务端的证书和KEY

	certPool := x509.NewCertPool()          //创建证书池
	ca, _ := ioutil.ReadFile("cert/ca.pem") //读取CA证书文件
	certPool.AppendCertsFromPEM(ca)         //把ca证书放入证书池

	creds := credentials.NewTLS(&tls.Config{ //grpc的传入对象
		Certificates: []tls.Certificate{cert},        //服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, //表明，需要验证客户端的证书，也就双向验证
		ClientCAs:    certPool,                       //指明客户端CA，是certPool

	})
	return creds
}

//获取Client端证书，供httpserver调用
func GetClientCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key") //读服客户端证书和KEY

	certPool := x509.NewCertPool()          //创建证书池
	ca, _ := ioutil.ReadFile("cert/ca.pem") //读取CA证书文件
	certPool.AppendCertsFromPEM(ca)         //把ca证书放入证书池

	creds := credentials.NewTLS(&tls.Config{ //grpc的传入对象
		Certificates: []tls.Certificate{cert}, //客户端证书
		ServerName:   "common",                //指定值，与证书中的域名相同
		RootCAs:      certPool,
	})
	return creds
}
