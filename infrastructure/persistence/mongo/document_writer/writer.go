package document_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type documentWriter struct {
	db *mongo.Client
}

// Constructor for documentWriter
func New(db *mongo.Client) repository.DocumentWriter {
	return &documentWriter{db: db}
}
