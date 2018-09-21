package database_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Creates a new "DeleteMe" collection to create DB.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
func (rcv *databaseWriter) Create(name value.DBName) error {
	return nil
}
