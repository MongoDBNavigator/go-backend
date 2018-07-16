package document_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"gopkg.in/mgo.v2"
)

type documentWriter struct {
	db *mgo.Session
}

// Constructor for documentWriter
func New(db *mgo.Session) repository.DocumentWriter {
	return &documentWriter{db: db}
}
