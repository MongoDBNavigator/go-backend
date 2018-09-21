package document_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Update document
// https://docs.mongodb.com/manual/tutorial/update-documents/
func (rcv *documentWriter) Update(dbName value.DBName, collName value.CollName, docId value.DocId, doc interface{}) error {
	return nil
}
