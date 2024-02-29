package main

import (
	config "go-exam/product-service/config"
	pb "go-exam/product-service/genproto/product"
	"go-exam/product-service/pkg/db"
	"go-exam/product-service/pkg/logger"
	service "go-exam/product-service/service"
	grpcClient "go-exam/product-service/service/grpc_client"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "prodservice")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatasbase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	grpcClien, err := grpcClient.New(cfg)

	if err != nil {
		log.Fatal("grpc client dial error", logger.Error(err))
	}

	productService := service.NewPostService(connDB, log, grpcClien)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, productService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
