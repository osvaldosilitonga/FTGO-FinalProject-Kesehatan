package main

import (
	"fmt"
	"log"
	"net"
	"order/configs"
	"order/controllers"
	"order/repository"
	"os"

	orderPB "order/internal/order"

	// grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

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

	orderController := controllers.NewOrderController(orderRepository)

	opts := []grpc.ServerOption{
		// grpc.ChainUnaryInterceptor(
		// 	logging.UnaryServerInterceptor(middlewares.NewInterceptorLogger()),
		// 	// grpc_auth.UnaryServerInterceptor(middlewares.JWTAuth),
		// ),
	}
	grpcServer := grpc.NewServer(opts...)
	orderPB.RegisterOrderServiceServer(grpcServer, orderController)

	log.Printf("Starting Order Service gRPC listener on port : %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
