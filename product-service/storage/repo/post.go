package repo

import (
	pb "go-exam/product-service/genproto/product"
)

// PostStorageI ...
type ProductStorageI interface {
	Create(pb *pb.Product) (*pb.Product, error)
	Get(req *pb.GetProductRequest) (*pb.Product, error)
	GetAll(*pb.GetAllRequest) (*pb.GetAllResponse, error)
	Update(*pb.Product) (*pb.Product, error)
	Delete(*pb.GetProductRequest) (*pb.Product, error)
}
