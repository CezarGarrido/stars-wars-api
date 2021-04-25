package infra

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	ctx, cancelFunc := context.WithTimeout(ctx, 12*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
