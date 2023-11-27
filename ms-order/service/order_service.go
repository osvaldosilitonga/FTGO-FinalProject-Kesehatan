package service

import (
	"context"
	pb "order/internal/order"

	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	order *pb.Order
	pb.UnimplementedOrderServiceServer
	collection *mongo.Collection
	client     *mongo.Client
}

func NewOrderService(col *mongo.Collection, cli *mongo.Client) *Order {
	return &Order{
		collection: col,
		client:     cli,
	}
}

func (o *Order) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	return &pb.Order{}, nil
}

func (o *Order) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	return &pb.Order{}, nil
}

func (o *Order) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return &pb.ListOrdersResponse{}, nil
}

func (o *Order) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.Order, error) {
	return &pb.Order{}, nil
}

func (o *Order) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.Order, error) {
	return &pb.Order{}, nil
}
