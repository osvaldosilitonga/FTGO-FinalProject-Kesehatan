package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "gateway/internal/product"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Product interface {
	ListProduct(ctx context.Context) (*pb.ListProductResponse, error)
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error)
	GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error)
	UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error)
	DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error)
	CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*emptypb.Empty, error)
	CheckProductExist(ctx context.Context, req *pb.CheckProductExistRequest) (*emptypb.Empty, error)
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

func (p *ProductImpl) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	productClient := pb.NewProductServiceClient(p.Conn)

	product, err := productClient.GetProduct(ctx, req)
	if err != nil {
		log.Printf("Error from get product service, err: %v\n", err)
		return nil, err
	}

	ca, err := strconv.ParseInt(product.CreatedAt, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parse created at, err: %v", err)
	}

	t := time.UnixMilli(ca)
	fmt.Println(t, "<----- created at")

	return product, nil
}

func (p *ProductImpl) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	productClient := pb.NewProductServiceClient(p.Conn)

	product, err := productClient.UpdateProduct(ctx, req)
	if err != nil {
		log.Printf("Error from update product service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error) {
	productClient := pb.NewProductServiceClient(p.Conn)

	product, err := productClient.DeleteProduct(ctx, req)
	if err != nil {
		log.Printf("Error from delete product service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*emptypb.Empty, error) {
	productClient := pb.NewProductServiceClient(p.Conn)

	product, err := productClient.CheckStock(ctx, req)
	if err != nil {
		log.Printf("Error from check stock service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) CheckProductExist(ctx context.Context, req *pb.CheckProductExistRequest) (*emptypb.Empty, error) {
	productClient := pb.NewProductServiceClient(p.Conn)

	product, err := productClient.CheckProductExist(ctx, req)
	if err != nil {
		log.Printf("Error from check product exist service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}
