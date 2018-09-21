package document_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Drop document
// https://docs.mongodb.com/manual/tutorial/remove-documents/
func (rcv *documentWriter) Delete(dbName value.DBName, collName value.CollName, docId value.DocId) error {
	return nil
}
