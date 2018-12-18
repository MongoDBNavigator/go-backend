package response

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
)

func DatabaseToView(database *model.Database) *representation.Database {
	return &representation.Database{
		Name:          database.Name(),
		Collections:   database.CollectionsNumber(),
		IndexesNumber: database.IndexesNumber(),
		StorageSize:   database.StorageSize(),
	}
}

func DatabasesToView(databases []*model.Database) *representation.Databases {
	view := new(representation.Databases)

	view.Objects = make([]*representation.Database, len(databases))

	for i, database := range databases {
		view.Objects[i] = DatabaseToView(database)
	}

	return view
}
