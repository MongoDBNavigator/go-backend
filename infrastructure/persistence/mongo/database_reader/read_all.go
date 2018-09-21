package database_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
)

// Returns a document that lists all databases and returns basic database statistics.
// https://docs.mongodb.com/manual/reference/command/listDatabases/
func (rcv *databaseReader) ReadAll() ([]*model.Database, error) {
	return nil, nil
}
