package database_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type databaseWriter struct {
	db *mongo.Client
}

// Constructor for databaseWriter
func New(db *mongo.Client) repository.DatabaseWriter {
	return &databaseWriter{db: db}
}
