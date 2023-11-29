package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBCollection(col string, client *mongo.Client) *mongo.Collection {
	return client.Database("final-project").Collection(col)
}

func InitDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_DB_STRING")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Println("Mongo DB connected Successfully")

	return client
}
