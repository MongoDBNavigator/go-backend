package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
)

func IndexesToView(models []*model.Index) *representation.Indexes {
	view := new(representation.Indexes)

	view.Objects = make([]*representation.Index, len(models))

	for i, index := range models {
		view.Objects[i] = IndexToView(index)
	}

	return view
}
