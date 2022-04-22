package services

import context "context"

type ProdService struct {
	UnimplementedProdServiceServer
}

/*
实现 _grpc.pb.go中的 ProdServiceServer 接口中的方法
*/
func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 25}, nil
}
