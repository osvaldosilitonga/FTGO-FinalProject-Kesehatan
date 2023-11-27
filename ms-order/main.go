package main

import (
	"fmt"
	"net"
	"os"

	"order/config"
	pb "order/internal/order"
	"order/service"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load .env file")
	}
}

func main() {
	port := os.Getenv("GRPC_SERVER_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	// MongoDB
	orderClient := config.InitDB()
	orderCollection := config.DBCollection("orders", orderClient)

	orderService := service.NewOrderService(orderCollection, orderClient)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderService)

	log.Printf("Starting gRPC listener on port : %v, network : %v", port, lis.Addr().Network())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
