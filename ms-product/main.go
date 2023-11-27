package main

import (
	"fmt"
	"log"
	"net"
	"product/configs"
	"product/service"

	pb "product/internal/product"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := "50051"
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	dbClient := configs.InitDB()
	defer dbClient.Disconnect(nil)
	dbCollection := configs.DBCollection("products", dbClient)

	productService := service.NewProductService(dbCollection, dbClient)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProductServiceServer(grpcServer, productService)

	log.Printf("Starting gRPC listener on port : %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
