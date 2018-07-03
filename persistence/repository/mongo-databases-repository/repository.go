package mongo_databases_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"gopkg.in/mgo.v2"
)

type databasesRepository struct {
	db *mgo.Session
}

func New(db *mgo.Session) repository.DatabasesRepositoryInterface {
	return &databasesRepository{
		db: db,
	}
}
