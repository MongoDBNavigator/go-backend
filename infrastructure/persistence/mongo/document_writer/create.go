package document_writer

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Insert new document
// https://docs.mongodb.com/manual/tutorial/insert-documents/
func (rcv *documentWriter) Create(dbName value.DBName, collName value.CollName, doc interface{}) error {
	return nil
}
