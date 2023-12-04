package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Products struct {
	Id    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Qty   int32  `json:"qty,omitempty" bson:"qty,omitempty"`
	Price int64  `json:"price,omitempty" bson:"price,omitempty"`
}

type Orders struct {
	Id          primitive.ObjectID `json:"order_id,omitempty" bson:"_id,omitempty"`
	UserId      int                `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty"`
	TotalAmount int64              `json:"total_amount,omitempty" bson:"total_amount,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Products    []Products         `json:"products,omitempty" bson:"products,omitempty"`
	CreatedAt   int64              `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
