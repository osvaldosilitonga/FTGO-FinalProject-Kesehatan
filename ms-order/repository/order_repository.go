package repository

import (
	"context"
	"errors"
	"order/models/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository interface {
	Save(ctx context.Context, data *entity.Orders) (*entity.Orders, error)
	FindById(ctx context.Context, id primitive.ObjectID) (*entity.Orders, error)
	Update(ctx context.Context, data *entity.Orders) error
	FindAll(ctx context.Context, page int, status string) ([]*entity.Orders, error)
	FindByUserID(ctx context.Context, userID, page int, status string) ([]*entity.Orders, error)
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

func (o *OrderRepositoryImpl) FindAll(ctx context.Context, page int, status string) ([]*entity.Orders, error) {
	var orders []*entity.Orders

	if page < 1 {
		page = 1
	}

	skip := int64((page - 1) * 10)
	limit := int64(10)

	filter := bson.M{}
	if status != "" {
		filter = bson.M{
			"status": status,
		}
	}

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: bson.M{
			"updated_at": -1,
		},
	}

	cursor, err := o.dbCollection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepositoryImpl) FindByUserID(ctx context.Context, userID, page int, status string) ([]*entity.Orders, error) {
	var orders []*entity.Orders

	if page < 1 {
		page = 1
	}

	skip := int64((page - 1) * 10)
	limit := int64(10)

	filter := bson.M{
		"user_id": userID,
	}
	if status != "" {
		filter = bson.M{
			"user_id": userID,
			"status":  status,
		}
	}

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: bson.M{
			"updated_at": -1,
		},
	}

	cursor, err := o.dbCollection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
