package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
)

func DatabasesToView(databases []*model.Database) *representation.Databases {
	view := new(representation.Databases)

	view.Objects = make([]*representation.Database, len(databases))

	for i, database := range databases {
		view.Objects[i] = DatabaseToView(database)
	}

	return view
}
