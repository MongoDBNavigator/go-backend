package index_writer

import (
	"gopkg.in/mgo.v2"

	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
)

type indexWriter struct {
	db *mgo.Session
}

// Constructor for indexWriter
func New(db *mgo.Session) repository.IndexWriter {
	return &indexWriter{db: db}
}
