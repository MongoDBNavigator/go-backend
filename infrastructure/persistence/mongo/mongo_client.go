package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// https://www.compose.com/articles/mongodb-and-go-moving-on-from-mgo/
func NewMongoDBClient(url string) (*mongo.Client, error) {
	client, err := mongo.NewClient(url)

	if err != nil {
		return nil, err
	}

	if err = client.Connect(context.Background()); err != nil {
		return nil, err
	}

	return client, err
}
