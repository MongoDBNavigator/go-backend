package validator_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/repository"
	"gopkg.in/mgo.v2"
)

type validatorWriter struct {
	db *mgo.Session
}

// Constructor for indexWriter
func New(db *mgo.Session) repository.ValidationWriter {
	return &validatorWriter{db: db}
}
