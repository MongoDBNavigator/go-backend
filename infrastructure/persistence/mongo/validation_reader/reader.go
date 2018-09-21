package validation_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type validationReader struct {
	db *mongo.Client
}

// Constructor for systemInfoReader
func New(db *mongo.Client) repository.ValidationReader {
	return &validationReader{
		db: db,
	}
}
