package service

import (
	"context"
	"fmt"

	pb "gateway/internal/order"

	"google.golang.org/grpc"
)

type Order interface {
	CreateOrderProduct(ctx context.Context, req *pb.CreateOrderProductRequest) (*pb.Order, error)
}

type OrderImpl struct {
	Conn *grpc.ClientConn
}

func NewOrderService(conn *grpc.ClientConn) Order {
	return &OrderImpl{
		Conn: conn,
	}
}

func (o *OrderImpl) CreateOrderProduct(ctx context.Context, req *pb.CreateOrderProductRequest) (*pb.Order, error) {
	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.CreateOrderProduct(ctx, req)
	if err != nil {
		fmt.Printf("----> Error from create order service, err: %v\n", err)
		return nil, err
	}

	return order, nil
}
