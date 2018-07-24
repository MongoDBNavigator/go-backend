package repository

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Interface for validation reader
// https://martinfowler.com/eaaCatalog/repository.html
type ValidationReader interface {
	// Fetch all validation info
	Read(dbName value.DBName, collName value.CollName) (*model.Validation, error)
}
