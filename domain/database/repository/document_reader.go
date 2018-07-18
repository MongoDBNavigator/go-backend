package repository

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Interface for document reader
// https://martinfowler.com/eaaCatalog/repository.html
type DocumentReader interface {
	// Fetch document by ID
	Read(dbName value.DBName, collName value.CollName, docId value.DocId) (interface{}, error)
	// Fetch documents with pagination and filters
	ReadAll(conditions *value.ReadAllDocConditions) ([]interface{}, error)
	// Get documents count by filters
	ReadCount(conditions *value.ReadAllDocConditions) (int, error)
}
