package repository

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Interface for database writer
// https://martinfowler.com/eaaCatalog/repository.html
type DatabaseWriter interface {
	// Drop database
	Delete(name value.DBName) error
	// Create database & "DeleteMe" collection
	Create(name value.DBName) error
}
