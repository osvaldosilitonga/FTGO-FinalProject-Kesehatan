package models

import (
	"order/internal/order"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type        string             `json:"type" bson:"type"`
	CustomerID  int32              `json:"customer_id" bson:"customer_id,omitempty"`
	Product     []*order.Product          `json:"product,omitempty" bson:"product,omitempty"`
	TotalAmount float64            `json:"total_amount" bson:"total_amount"`
	CreatedAt   string             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   string             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Product struct {
	Id          string `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Price       int64  `json:"price,omitempty" bson:"price,omitempty"`
	Quantity    int    `json:"qty,omitempty" bson:"qty,omitempty"`
}
