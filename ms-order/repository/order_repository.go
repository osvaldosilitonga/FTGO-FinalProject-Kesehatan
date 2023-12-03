package repository

import (
	"context"
	"order/models/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Save(ctx context.Context, data *entity.Orders) (*entity.Orders, error)
}

type OrderRepositoryImpl struct {
	dbCollection *mongo.Collection
	dbClient     *mongo.Client
}

func NewOrderRepository(col *mongo.Collection, cli *mongo.Client) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		dbCollection: col,
		dbClient:     cli,
	}
}

func (o *OrderRepositoryImpl) Save(ctx context.Context, data *entity.Orders) (*entity.Orders, error) {
	result, err := o.dbCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	data.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return data, nil
}
