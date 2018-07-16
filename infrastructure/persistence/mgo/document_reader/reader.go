package document_reader

import (
	"gopkg.in/mgo.v2"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
)

type documentReader struct {
	db *mgo.Session
}

// Constructor for documentReader
func New(db *mgo.Session) repository.DocumentReader {
	return &documentReader{db: db}
}
