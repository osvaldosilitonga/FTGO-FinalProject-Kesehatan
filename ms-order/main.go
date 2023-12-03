package main

import (
	"fmt"
	"log"
	"net"
	"order/configs"
	"order/repository"
	"order/services"
	"os"

	orderPB "order/internal/order"

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
	dbCollection := configs.DBCollection("orders", dbClient)

	orderRepository := repository.NewOrderRepository(dbCollection, dbClient)

	// productGrpc := configs.ProductGrpc(os.Getenv("PRODUCT_GRPC_SERVER"))
	// productService := services.NewProductService(productGrpc)

	orderService := services.NewOrderService(orderRepository)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	orderPB.RegisterOrderServiceServer(grpcServer, orderService)

	log.Printf("Starting Order Service gRPC listener on port : %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
