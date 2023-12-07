package service

import (
	"context"
	"fmt"
	"time"

	pb "gateway/internal/order"
	"gateway/middlewares"

	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
)

type Order interface {
	CreateOrderProduct(ctx context.Context, req *pb.CreateOrderProductRequest) (*pb.Order, error)
	FindByOrderId(ctx context.Context, req *pb.FindByOrderIdRequest) (*pb.Order, error)
	UpdateStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.Order, error)
	ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderResponse, error)
	FindByUserId(ctx context.Context, req *pb.FindByUserIdRequest) (*pb.ListOrderResponse, error)
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
	// GRPC Auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := middlewares.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.CreateOrderProduct(ctxWithAuth, req)

	// order, err := orderClient.CreateOrderProduct(ctx, req)
	if err != nil {
		fmt.Printf("----> Error from create order service, err: %v\n", err)
		return nil, err
	}

	return order, nil
}

func (o *OrderImpl) FindByOrderId(ctx context.Context, req *pb.FindByOrderIdRequest) (*pb.Order, error) {
	// GRPC Auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := middlewares.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.FindByOrderId(ctxWithAuth, req)
	// order, err := orderClient.FindByOrderId(ctx, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderImpl) UpdateStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.Order, error) {
	// GRPC Auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := middlewares.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.UpdateStatus(ctxWithAuth, req)
	// order, err := orderClient.UpdateStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderImpl) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
	// GRPC Auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := middlewares.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.ListOrder(ctxWithAuth, req)

	// order, err := orderClient.ListOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderImpl) FindByUserId(ctx context.Context, req *pb.FindByUserIdRequest) (*pb.ListOrderResponse, error) {
	// GRPC Auth
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := middlewares.SignJwtForGrpc()
	if err != nil {
		return nil, err
	}

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	orderClient := pb.NewOrderServiceClient(o.Conn)

	order, err := orderClient.FindByUserID(ctxWithAuth, req)

	// order, err := orderClient.FindByUserID(ctx, req)
	if err != nil {
		return nil, err
	}

	return order, nil
}
