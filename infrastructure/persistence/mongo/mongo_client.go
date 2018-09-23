package mongo

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func MongoDBClientFactory(url string) (*mongo.Client, error) {
	client, err := mongo.NewClient(url)

	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())

	if err != nil {
		return nil, err
	}

	return client, err
}
