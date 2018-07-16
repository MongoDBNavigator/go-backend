package repository

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Interface for collection writer
// https://martinfowler.com/eaaCatalog/repository.html
type CollectionWriter interface {
	// Drop collection
	Delete(dbName value.DBName, collName value.CollName) error
	// Create collection
	Create(dbName value.DBName, collName value.CollName) error
}
