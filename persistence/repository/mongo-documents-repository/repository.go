package mongo_documents_repository

import (
	"gopkg.in/mgo.v2"
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
)

type documentsRepository struct {
	db *mgo.Session
}

func New(db *mgo.Session) repository.DocumentsRepositoryInterface {
	return &documentsRepository{
		db: db,
	}
}
