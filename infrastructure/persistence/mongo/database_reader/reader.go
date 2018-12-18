package database_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Structure to implement DatabaseReader interface
type databaseReader struct {
	db *mongo.Client
}

// Constructor for databaseReader
func New(db *mongo.Client) repository.DatabaseReader {
	return &databaseReader{db: db}
}
