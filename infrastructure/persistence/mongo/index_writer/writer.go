package index_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Structure to implement IndexWriter interface
type indexWriter struct {
	db *mongo.Client
}

// Constructor for indexWriter
func New(db *mongo.Client) repository.IndexWriter {
	return &indexWriter{db: db}
}
