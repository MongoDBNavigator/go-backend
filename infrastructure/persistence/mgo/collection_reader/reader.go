package collection_reader

import (
	"gopkg.in/mgo.v2"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
)

type collectionReader struct {
	db *mgo.Session
}

// Constructor for collectionReader
func New(db *mgo.Session) repository.CollectionReader {
	return &collectionReader{db: db}
}
