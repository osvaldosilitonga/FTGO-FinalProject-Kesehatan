package service

import (
	"context"
	"product/entity"
	"product/helper"
	pb "product/internal/product"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Product struct {
	product *pb.Product
	pb.UnimplementedProductServiceServer
	dbCollection *mongo.Collection
}

func NewProductService(col *mongo.Collection) *Product {
	return &Product{
		dbCollection: col,
	}
}

func (p *Product) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	now := time.Now().UnixMilli()

	if req.Price <= 1000 {
		return nil, status.Error(codes.InvalidArgument, "price must be greater than 1000")
	}
	if req.Stock <= 1 {
		return nil, status.Error(codes.InvalidArgument, "stock must be greater than 1")
	}

	product := &entity.Products{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   strconv.Itoa(int(now)),
		UpdatedAt:   strconv.Itoa(int(now)),
	}

	result, err := p.dbCollection.InsertOne(ctx, product)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	product.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return helper.ToProductResponse(product), nil
}

func (p *Product) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	return &pb.Product{}, nil
}

func (p *Product) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	return &pb.Product{}, nil
}

func (p *Product) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error) {
	return &pb.Product{}, nil
}

func (p *Product) ListProduct(ctx context.Context, req *pb.Empty) (*pb.ListProductResponse, error) {
	return &pb.ListProductResponse{}, nil
}
