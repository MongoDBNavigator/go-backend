package database_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"gopkg.in/mgo.v2"
)

type databaseReader struct {
	db *mgo.Session
}

// Constructor for databaseReader
func New(db *mgo.Session) repository.DatabaseReader {
	return &databaseReader{db: db}
}
