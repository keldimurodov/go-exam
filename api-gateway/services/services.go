package services

import (
	"fmt"

	"go-exam/api-gateway/config"
	pbp "go-exam/api-gateway/genproto/product"
	pbu "go-exam/api-gateway/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	ProductService() pbp.ProductServiceClient
}

type serviceManager struct {
	userService    pbu.UserServiceClient
	productService pbp.ProductServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) ProductService() pbp.ProductServiceClient {
	return s.productService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ProductServiceHost, conf.ProductServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:    pbu.NewUserServiceClient(connUser),
		productService: pbp.NewProductServiceClient(connPost),
	}

	return serviceManager, nil
}
