package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
)

func IndexToView(model *model.Index) *representation.Index {
	return &representation.Index{
		Name:       model.Name(),
		Unique:     model.Unique(),
		Background: model.Background(),
		Sparse:     model.Sparse(),
		Fields:     model.Fields(),
	}
}
