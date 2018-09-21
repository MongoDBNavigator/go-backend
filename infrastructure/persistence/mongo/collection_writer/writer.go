package collection_writer

import (
	"gopkg.in/mgo.v2"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
)

type collectionWriter struct {
	db *mgo.Session
}

// Constructor for collectionWriter
func New(db *mgo.Session) repository.CollectionWriter {
	return &collectionWriter{db: db}
}
