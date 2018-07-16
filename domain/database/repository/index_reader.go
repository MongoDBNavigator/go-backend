package repository

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Interface for index reader
// https://martinfowler.com/eaaCatalog/repository.html
type IndexReader interface {
	// Fetch all indexes
	ReadAll(dbName value.DBName, collName value.CollName) ([]*model.Index, error)
}
