package repository

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Interface for validation writer
// https://martinfowler.com/eaaCatalog/repository.html
type ValidationWriter interface {
	// Add validator to collection
	Write(dbName value.DBName, collName value.CollName) error
}
