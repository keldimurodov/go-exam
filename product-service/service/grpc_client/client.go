package grpcClient

import (
	"fmt"
	config "go-exam/product-service/config"
	pbu "go-exam/product-service/genproto/user"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
}
type serviceManager struct {
	cfg             config.Config
	User pbu.UserServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.User

}

func New(cfg config.Config) (IServiceManager, error) {
	UserConnection, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ProductServiceHost, cfg.ProductServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &serviceManager{
		cfg:             cfg,
		User: pbu.NewUserServiceClient(UserConnection)}, nil

}
