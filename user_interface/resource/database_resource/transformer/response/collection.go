package response

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"
)

func CollectionToView(collection *model.Collection) *representation.Collection {
	return &representation.Collection{
		Name:            collection.Name(),
		DocumentsNumber: collection.DocNumber(),
		IndexesNumber:   collection.IndexesNumber(),
		AvgObjSize:      collection.AvgObjSize(),
	}
}

func CollectionsToView(collections []*model.Collection) *representation.Collections {
	view := new(representation.Collections)

	view.Objects = make([]*representation.Collection, len(collections))

	for i, collection := range collections {
		view.Objects[i] = CollectionToView(collection)
	}

	return view
}
