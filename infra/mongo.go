package infra

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	MONGO_DEFAULT_URL = "mongodb://localhost:27017"
)
func NewMongoClient(url string) (*mongo.Client, error) {

	clientOpts := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
