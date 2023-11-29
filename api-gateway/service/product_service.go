package service

import (
	"context"
	"log"

	pb "gateway/internal/product"

	"google.golang.org/grpc"
)

type Product interface {
	ListProduct(ctx context.Context) (*pb.ListProductResponse, error)
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error)
}

type ProductImpl struct {
	Conn *grpc.ClientConn
}

func NewProductService(conn *grpc.ClientConn) *ProductImpl {
	return &ProductImpl{
		Conn: conn,
	}
}

func (p *ProductImpl) ListProduct(ctx context.Context) (*pb.ListProductResponse, error) {
	productClient := pb.NewProductServiceClient(p.Conn)

	list, err := productClient.ListProduct(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("Error from list product service, err: %v\n", err)
		return nil, err
	}

	return list, nil
}

func (p *ProductImpl) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	productClient := pb.NewProductServiceClient(p.Conn)

	product, err := productClient.CreateProduct(ctx, req)
	if err != nil {
		log.Printf("Error from create product service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}
