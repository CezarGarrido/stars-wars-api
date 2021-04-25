package infra

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_DEFAULT_URL = "mongodb://localhost:27017"
)

//NewMongoClient Returns a new mongo Client.
//There is a timeout of 12 s to initiate the connection.
//Returns error if connection to MongoDB fails.
func NewMongoClient(url string) (*mongo.Client, error) {
	ctx := context.TODO()

	clientOpts := options.Client().ApplyURI(url)

	ctx, cancelFunc := context.WithTimeout(ctx, 12000)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
