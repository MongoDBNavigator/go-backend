package repository

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Interface for collection reader
// https://martinfowler.com/eaaCatalog/repository.html
type CollectionReader interface {
	// Fetch all collections in database
	ReadAll(dbName value.DBName) ([]*model.Collection, error)
}
