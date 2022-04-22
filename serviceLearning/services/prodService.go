package services

import context "context"

/*
实现ProdServiceServer接口
*/
type ProdService struct {
	UnimplementedProdServiceServer //_grpc_pb.go中提供的默认实现
}

/*
实现 _grpc.pb.go中的 ProdServiceServer 接口中的方法
*/
func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 25}, nil
}
