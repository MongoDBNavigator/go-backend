package index_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"

	"gopkg.in/mgo.v2"
)

type indexReader struct {
	db *mgo.Session
}

// Constructor for indexReader
func New(db *mgo.Session) repository.IndexReader {
	return &indexReader{db: db}
}
