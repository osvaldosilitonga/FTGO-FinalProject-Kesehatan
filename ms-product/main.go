package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"product/configs"
	"product/middlewares"
	"product/service"

	// "product/service"

	pb "product/internal/product"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PRODUCT_SERVICE_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	dbClient := configs.InitDB()
	defer dbClient.Disconnect(nil)
	dbCollection := configs.DBCollection("products", dbClient)

	productService := service.NewProductService(dbCollection, dbClient)

	// var opts []grpc.ServerOption
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middlewares.NewInterceptorLogger()),
			grpc_auth.UnaryServerInterceptor(middlewares.JWTAuth),
		),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductServiceServer(grpcServer, productService)

	log.Printf("Starting gRPC listener on port : %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
