package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
)

func CollectionsToView(collections []*model.Collection) *representation.Collections {
	view := new(representation.Collections)

	view.Objects = make([]*representation.Collection, len(collections))

	for i, collection := range collections {
		view.Objects[i] = CollectionToView(collection)
	}

	return view
}
