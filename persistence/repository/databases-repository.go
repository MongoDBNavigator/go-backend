package repository

import "github.com/MongoDBNavigator/go-backend/persistence/model"

type DatabasesRepositoryInterface interface {
	GetListDatabases() ([]*model.Database, error)
	DropDatabase(name string) error
	RawQuery(jsonQuery string) (interface{}, error)
}
