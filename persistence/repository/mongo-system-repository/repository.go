package mongo_system_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"gopkg.in/mgo.v2"
)

type systemRepository struct {
	db  *mgo.Session
	url string
}

func New(db *mgo.Session, url string) repository.SystemRepositoryInterface {
	return &systemRepository{
		db:  db,
		url: url,
	}
}
