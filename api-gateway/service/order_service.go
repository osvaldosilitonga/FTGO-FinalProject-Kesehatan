package service

import (
	"context"
	"fmt"

	pb "gateway/internal/order"

	"google.golang.org/grpc"
)

type Order interface {
	CreateOrderProduct(ctx context.Context, req *pb.CreateOrderProductRequest) (*pb.Order, error)
	FindByOrderId(ctx context.Context, req *pb.FindByOrderIdRequest) (*pb.Order, error)
	UpdateStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.Order, error)
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

func (o *OrderImpl) FindByOrderId(ctx context.Context, req *pb.FindByOrderIdRequest) (*pb.Order, error) {
	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.FindByOrderId(ctx, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderImpl) UpdateStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.Order, error) {
	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.UpdateStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}
