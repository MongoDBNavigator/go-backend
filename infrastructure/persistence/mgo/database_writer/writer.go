package database_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"gopkg.in/mgo.v2"
)

type databaseWriter struct {
	db *mgo.Session
}

// Constructor for databaseWriter
func New(db *mgo.Session) repository.DatabaseWriter {
	return &databaseWriter{db: db}
}
