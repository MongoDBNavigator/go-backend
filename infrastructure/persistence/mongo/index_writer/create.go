package index_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Creates indexes on collections.
// https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
func (rcv *indexWriter) Create(dbName value.DBName, collName value.CollName, index *model.Index) error {
	return nil
}
