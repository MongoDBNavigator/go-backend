package validator_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Structure to implement ValidationWriter interface
type validatorWriter struct {
	db *mongo.Client
}

// Constructor for indexWriter
func New(db *mongo.Client) repository.ValidationWriter {
	return &validatorWriter{db: db}
}
