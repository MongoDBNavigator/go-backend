package document_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Fetch document by ID
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) Read(dbName value.DBName, collName value.CollName, docId value.DocId) (interface{}, error) {
	return nil, nil
}
