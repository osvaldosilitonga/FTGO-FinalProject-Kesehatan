package helper

import (
	pb "order/internal/order"
	"order/models/entity"
)

func ToOrderResponse(data *entity.Orders) *pb.Order {

	products := []*pb.Product{}
	for _, product := range data.Products {
		p := pb.Product{}
		p.Id = product.Id
		p.Qty = product.Qty

		products = append(products, &p)
	}

	return &pb.Order{
		OrderId:     data.Id,
		UserId:      int32(data.UserId),
		Type:        data.Type,
		TotalAmount: data.TotalAmount,
		Status:      data.Status,
		Products:    products,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}
