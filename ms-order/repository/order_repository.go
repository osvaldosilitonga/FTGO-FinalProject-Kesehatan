package repository

import (
	"context"
	"errors"
	"order/models/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	Save(ctx context.Context, data *entity.Orders) (*entity.Orders, error)
	FindById(ctx context.Context, id primitive.ObjectID) (*entity.Orders, error)
	Update(ctx context.Context, data *entity.Orders) error
}

type OrderRepositoryImpl struct {
	dbCollection *mongo.Collection
	dbClient     *mongo.Client
}

func NewOrderRepository(col *mongo.Collection, cli *mongo.Client) OrderRepository {
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

	data.Id = result.InsertedID.(primitive.ObjectID)

	return data, nil
}

func (o *OrderRepositoryImpl) Update(ctx context.Context, data *entity.Orders) error {
	result := o.dbCollection.FindOneAndUpdate(ctx, bson.M{"_id": data.Id}, bson.M{"$set": data})
	if result.Err() == mongo.ErrNoDocuments {
		return errors.New("order not found")
	}
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (o *OrderRepositoryImpl) FindById(ctx context.Context, id primitive.ObjectID) (*entity.Orders, error) {
	var order entity.Orders

	result := o.dbCollection.FindOne(ctx, bson.M{"_id": id})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, errors.New("order not found")
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
