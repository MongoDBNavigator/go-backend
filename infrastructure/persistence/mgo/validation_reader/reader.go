package validation_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"gopkg.in/mgo.v2"
)

type validationReader struct {
	db *mgo.Session
}

// Constructor for systemInfoReader
func New(db *mgo.Session) repository.ValidationReader {
	return &validationReader{
		db: db,
	}
}
