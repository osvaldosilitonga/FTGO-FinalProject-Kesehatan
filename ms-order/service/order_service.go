package service

import (
	"context"
	"fmt"
	"log"
	"order/config"
	pb "order/internal/order"
	model "order/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order struct {
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

type CreateOrderReturnType struct {
	Id string
	model.Order
}

func (o *Order) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	// Validasi apakah request body sesuai dengan kebutuhan
	for _, product := range req.Product {
		if product.Id == "" || product.Quantity <= 0 {
			err := status.Errorf(codes.InvalidArgument, "Invalid request body")
			log.Printf("Invalid request body: %v", err)
			return nil, err
		}
	}

	productClient := config.InitDB()
	productCollection := config.DBCollection("products", productClient)

	// Lakukan validasi atau logika bisnis lainnya sebelum menyimpan order ke database
	// Misalnya, pastikan produk dengan ID yang diberikan tersedia

	// Contoh: Validasi ketersediaan produk

	var totalAmount int32

	for _, product := range req.Product {
		totalAmount += product.Quantity * int32(product.Price)
		productId, _ := primitive.ObjectIDFromHex(product.Id)
		foundProduct := productCollection.FindOne(ctx, bson.M{"_id": productId})
		fmt.Printf("%s", product.Id)
		if foundProduct.Err() != nil {
			log.Printf("Product not available: %v", foundProduct.Err())
			return nil, foundProduct.Err()
		}
	}

	newOrder := pb.Order{
		// Isi data order dari request sesuai kebutuhan proyek Anda
		Type:        req.Type,
		CustomerId:  req.CustomerId,
		Product:     req.Product,
		TotalAmount: int64(totalAmount),
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
		// Tambahan data order lainnya
	}

	// Simpan order ke database
	insertedResult, err := o.collection.InsertOne(ctx, &newOrder)
	if err != nil {
		log.Printf("Error inserting order: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create order")
	}

	// Setel ID hasil penyisipan ke order yang baru dibuat

	newOrder.Id = insertedResult.InsertedID.(primitive.ObjectID).Hex()

	return &newOrder, nil
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
