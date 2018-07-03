package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
)

func DatabaseToView(database *model.Database) *representation.Database {
	return &representation.Database{
		Name:          database.GetName(),
		Collections:   database.GetCollectionsNumber(),
		IndexesNumber: database.GetIndexesNumber(),
		StorageSize:   database.GetStorageSize(),
	}
}
