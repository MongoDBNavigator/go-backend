package repository

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Interface for document writer
// https://martinfowler.com/eaaCatalog/repository.html
type DocumentWriter interface {
	// Insert document
	Create(dbName value.DBName, collName value.CollName, doc interface{}) error
	// Drop document
	Delete(dbName value.DBName, collName value.CollName, docId value.DocId) error
	// Update document
	Update(dbName value.DBName, collName value.CollName, docId value.DocId, doc interface{}) error
}
