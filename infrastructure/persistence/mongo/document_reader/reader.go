package document_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type documentReader struct {
	db *mongo.Client
}

// Constructor for documentReader
func New(db *mongo.Client) repository.DocumentReader {
	return &documentReader{db: db}
}
