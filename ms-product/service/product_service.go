package service

import (
	"context"
	"errors"
	"fmt"
	"product/entity"
	"product/helper"
	pb "product/internal/product"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Product struct {
	product *pb.Product
	pb.UnimplementedProductServiceServer
	dbCollection *mongo.Collection
	dbClient     *mongo.Client
}

func NewProductService(col *mongo.Collection, cli *mongo.Client) *Product {
	return &Product{
		dbCollection: col,
		dbClient:     cli,
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
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	product := &entity.Products{}
	filter := bson.M{"_id": id}
	err = p.dbCollection.FindOne(ctx, filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return helper.ToProductResponse(product), nil
}

func (p *Product) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	now := time.Now().UnixMilli()
	product := &entity.Products{
		UpdatedAt: strconv.Itoa(int(now)),
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Category != "" {
		product.Category = req.Category
	}
	if req.Price >= 1000 {
		product.Price = req.Price
	}
	if req.Stock >= 1 {
		product.Stock = req.Stock
	}

	result := p.dbCollection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": product})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = p.dbCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return helper.ToProductResponse(product), nil
}

func (p *Product) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Product, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	product := &entity.Products{}
	err = p.dbCollection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return helper.ToProductResponse(product), nil
}

func (p *Product) ListProduct(ctx context.Context, req *pb.Empty) (*pb.ListProductResponse, error) {

	products := []entity.Products{}
	cursor, err := p.dbCollection.Find(ctx, bson.M{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("product not found"))
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		product := entity.Products{}
		if err := cursor.Decode(&product); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		products = append(products, product)
	}

	pr := []*pb.Product{}
	for _, product := range products {
		p := helper.ToProductResponse(&product)
		pr = append(pr, p)
	}

	productsResponse := &pb.ListProductResponse{
		Products: pr,
	}

	return productsResponse, nil
}

func (p *Product) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*emptypb.Empty, error) {

	for _, product := range req.Datas {
		id, err := primitive.ObjectIDFromHex(product.Id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		data := &entity.Products{}

		filter := bson.M{"_id": id}
		err = p.dbCollection.FindOne(ctx, filter).Decode(&data)
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("product with id: %s, not found", product.Id))
		}
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if data.Stock < product.Quantity {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("product with id: %s, out of stock", product.Id))
		}
	}

	return new(emptypb.Empty), nil
}

func (p *Product) CheckProductExist(ctx context.Context, req *pb.CheckProductExistRequest) (*emptypb.Empty, error) {

	for _, product := range req.Datas {
		id, err := primitive.ObjectIDFromHex(product)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		filter := bson.M{"_id": id}
		err = p.dbCollection.FindOne(ctx, filter).Err()
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("product with id: %s, not found", product))
		}
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

	}

	return new(emptypb.Empty), nil
}
