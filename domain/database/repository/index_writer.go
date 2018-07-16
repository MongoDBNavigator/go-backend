package repository

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Interface for index writer
// https://martinfowler.com/eaaCatalog/repository.html
type IndexWriter interface {
	// Method for dropping index in collection
	Delete(dbName value.DBName, collName value.CollName, indexName value.IndexName) error
	// Method for creating index in collection
	Create(dbName value.DBName, collName value.CollName, index *model.Index) error
}
