package collection_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"gopkg.in/mgo.v2"
)

type collectionWriter struct {
	db *mgo.Session
}

// Constructor for collectionWriter
func New(db *mgo.Session) repository.CollectionWriter {
	return &collectionWriter{db: db}
}
