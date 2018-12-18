package collection_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Structure to implement CollectionWriter interface
type collectionWriter struct {
	db *mongo.Client
}

// Constructor for collectionWriter
func New(db *mongo.Client) repository.CollectionWriter {
	return &collectionWriter{db: db}
}
