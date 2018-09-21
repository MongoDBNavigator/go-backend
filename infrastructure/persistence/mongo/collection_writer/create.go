package collection_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Creates a new collection or view.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
func (rcv *collectionWriter) Create(dbName value.DBName, collName value.CollName) error {
	return nil
}
