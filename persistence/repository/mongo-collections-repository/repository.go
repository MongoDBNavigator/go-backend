package mongo_collections_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"gopkg.in/mgo.v2"
)

type collectionsRepository struct {
	db *mgo.Session
}

func New(db *mgo.Session) repository.CollectionsRepositoryInterface {
	return &collectionsRepository{
		db: db,
	}
}
