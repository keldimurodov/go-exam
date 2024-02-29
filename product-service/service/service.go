package service

import (
	"context"
	pb "go-exam/product-service/genproto/product"
	l "go-exam/product-service/pkg/logger"
	grpcClient "go-exam/product-service/service/grpc_client"
	storage "go-exam/product-service/storage"

	"github.com/jmoiron/sqlx"
)

// PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// Delete implements product.ProductServiceServer.
func (s *PostService) Delete(context.Context, *pb.GetProductRequest) (*pb.Product, error) {
	panic("unimplemented")
}

// Update implements product.ProductServiceServer.
func (s *PostService) Update(context.Context, *pb.Product) (*pb.Product, error) {
	panic("unimplemented")
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	pro, err := s.storage.Product().Create(req)
	if err != nil {
		return nil, err
	}
	return pro, nil
}

func (s *PostService) Get(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	pro, err := s.storage.Product().Get(req)
	if err != nil {
		s.logger.Error(err.Error())
	}

	return pro, nil
}

func (s *PostService) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	med, err := s.storage.Product().GetAll(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return med, nil
}
