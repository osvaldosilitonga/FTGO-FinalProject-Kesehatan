package helper

import (
	"product/entity"
	pb "product/internal/product"
)

func ToProductResponse(data *entity.Products) *pb.Product {
	return &pb.Product{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Category:    data.Category,
		Price:       data.Price,
		Stock:       data.Stock,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}
