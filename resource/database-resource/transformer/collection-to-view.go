package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
)

func CollectionToView(collection *model.Collection) *representation.Collection {
	return &representation.Collection{
		Name:            collection.GetName(),
		DocumentsNumber: collection.GetDocumentsNumber(),
		IndexesNumber:   collection.GetIndexesNumber(),
		AvgObjSize:      collection.GetAvgObjSize(),
	}
}
