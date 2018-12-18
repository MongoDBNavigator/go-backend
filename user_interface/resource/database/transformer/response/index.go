package response

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
)

func IndexToView(model *model.Index) *representation.Index {
	return &representation.Index{
		Name:                    model.Name(),
		Unique:                  model.Unique(),
		Background:              model.Background(),
		Sparse:                  model.Sparse(),
		Fields:                  model.Fields(),
		PartialFilterExpression: model.PartialFilterExpression(),
	}
}

func IndexesToView(models []*model.Index) *representation.Indexes {
	view := new(representation.Indexes)

	view.Objects = make([]*representation.Index, len(models))

	for i, index := range models {
		view.Objects[i] = IndexToView(index)
	}

	return view
}
