package index_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type indexReader struct {
	db *mongo.Client
}

// Constructor for indexReader
func New(db *mongo.Client) repository.IndexReader {
	return &indexReader{db: db}
}
