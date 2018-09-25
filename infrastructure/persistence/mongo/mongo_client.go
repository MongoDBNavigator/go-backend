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

	if err = client.Connect(context.Background()); err != nil {
		return nil, err
	}

	return client, err
}
