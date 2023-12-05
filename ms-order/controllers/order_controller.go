package controllers

import (
	"context"
	"order/helper"
	orderPb "order/internal/order"
	productPb "order/internal/product"
	"order/models/entity"
	"order/repository"
	"order/services"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order struct {
	order *orderPb.Order
	orderPb.UnimplementedOrderServiceServer
	orderRepo repository.OrderRepository
}

func NewOrderController(or repository.OrderRepository) *Order {
	return &Order{
		orderRepo: or,
	}
}

func (o *Order) CreateOrderProduct(ctx context.Context, req *orderPb.CreateOrderProductRequest) (*orderPb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check stock
	data := &productPb.CheckStockRequest{}
	for _, product := range req.Products {
		d := &productPb.Data{}
		d.Id = product.Id
		d.Quantity = product.Qty

		data.Datas = append(data.Datas, d)
	}

	productsDetails, err := services.CheckStock(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// count total amount
	var totalAmount int64
	products := []entity.Products{}
	for _, product := range productsDetails.Products {
		for _, reqProduct := range req.Products {
			if product.Id == reqProduct.Id {
				totalAmount += int64(product.Price) * int64(reqProduct.Qty)

				p := entity.Products{}
				p.Id = product.Id
				p.Qty = reqProduct.Qty
				p.Name = product.Name
				p.Price = product.Price

				products = append(products, p)
			}
		}
	}

	now := time.Now().UnixMilli()
	order := &entity.Orders{
		UserId:      int(req.User.Id),
		UserEmail:   req.User.Email,
		Type:        "product",
		TotalAmount: int64(totalAmount),
		Status:      "PENDING",
		Products:    products,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	r, err := o.orderRepo.Save(ctx, order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := helper.ToOrderResponse(r)

	return response, nil
}

func (o *Order) UpdateStatus(ctx context.Context, req *orderPb.UpdateOrderStatusRequest) (*orderPb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	orderId, err := primitive.ObjectIDFromHex(req.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	order, err := o.orderRepo.FindById(ctx, orderId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	order.Status = strings.ToUpper(req.Status)
	order.UpdatedAt = time.Now().UnixMilli()

	err = o.orderRepo.Update(ctx, order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := helper.ToOrderResponse(order)

	return response, nil
}
