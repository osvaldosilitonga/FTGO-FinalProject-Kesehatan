package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "gateway/internal/product"
	// "gateway/middlewares"

	"google.golang.org/grpc"
	// grpcMetadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Product interface {
	ListProduct(ctx context.Context) (*pb.ListProductResponse, error)
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error)
	GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error)
	UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error)
	DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error)
	CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.ListProductResponse, error)
	CheckProductExist(ctx context.Context, req *pb.CheckProductExistRequest) (*emptypb.Empty, error)
	UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.ListProductResponse, error)
}

type ProductImpl struct {
	Conn *grpc.ClientConn
}

func NewProductService(conn *grpc.ClientConn) Product {
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
	// // GRPC Auth
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// token, err := middlewares.SignJwtForGrpc()
	// if err != nil {
	// 	return nil, err
	// }

	// ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	productClient := pb.NewProductServiceClient(p.Conn)

	// product, err := productClient.CreateProduct(ctxWithAuth, req)

	product, err := productClient.CreateProduct(ctx, req)
	if err != nil {
		log.Printf("Error from create product service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	// // GRPC Auth
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// token, err := middlewares.SignJwtForGrpc()
	// if err != nil {
	// 	return nil, err
	// }

	// ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	productClient := pb.NewProductServiceClient(p.Conn)

	// product, err := productClient.GetProduct(ctxWithAuth, req)
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
	// // GRPC Auth
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// token, err := middlewares.SignJwtForGrpc()
	// if err != nil {
	// 	return nil, err
	// }

	// ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	productClient := pb.NewProductServiceClient(p.Conn)

	// product, err := productClient.UpdateProduct(ctxWithAuth, req)
	product, err := productClient.UpdateProduct(ctx, req)
	if err != nil {
		log.Printf("Error from update product service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error) {
	// // GRPC Auth
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// token, err := middlewares.SignJwtForGrpc()
	// if err != nil {
	// 	return nil, err
	// }

	// ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	productClient := pb.NewProductServiceClient(p.Conn)

	// product, err := productClient.DeleteProduct(ctxWithAuth, req)
	product, err := productClient.DeleteProduct(ctx, req)
	if err != nil {
		log.Printf("Error from delete product service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.ListProductResponse, error) {
	// // GRPC Auth
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// token, err := middlewares.SignJwtForGrpc()
	// if err != nil {
	// 	return nil, err
	// }

	// ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	productClient := pb.NewProductServiceClient(p.Conn)

	// product, err := productClient.CheckStock(ctxWithAuth, req)
	product, err := productClient.CheckStock(ctx, req)
	if err != nil {
		log.Printf("Error from check stock service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) CheckProductExist(ctx context.Context, req *pb.CheckProductExistRequest) (*emptypb.Empty, error) {
	// // GRPC Auth
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// token, err := middlewares.SignJwtForGrpc()
	// if err != nil {
	// 	return nil, err
	// }

	// ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	productClient := pb.NewProductServiceClient(p.Conn)

	// product, err := productClient.CheckProductExist(ctxWithAuth, req)
	product, err := productClient.CheckProductExist(ctx, req)
	if err != nil {
		log.Printf("Error from check product exist service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}

func (p *ProductImpl) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.ListProductResponse, error) {
	// // GRPC Auth
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()

	// token, err := middlewares.SignJwtForGrpc()
	// if err != nil {
	// 	return nil, err
	// }

	// ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	productClient := pb.NewProductServiceClient(p.Conn)

	// product, err := productClient.UpdateStock(ctxWithAuth, req)
	product, err := productClient.UpdateStock(ctx, req)
	if err != nil {
		log.Printf("Error from update stock service, err: %v\n", err)
		return nil, err
	}

	return product, nil
}
