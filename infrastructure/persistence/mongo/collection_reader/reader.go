package collection_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type collectionReader struct {
	db *mongo.Client
}

// Constructor for collectionReader
func New(db *mongo.Client) repository.CollectionReader {
	return &collectionReader{db: db}
}
