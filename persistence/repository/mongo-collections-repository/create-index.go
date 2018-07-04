package mongo_collections_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"gopkg.in/mgo.v2"
)

//
// Creates indexes on collections.
// https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
//
func (rcv *collectionsRepository) CreateIndex(databaseName string, collectionName string, index *model.Index) error {
	i := mgo.Index{
		Name:       index.Name(),
		Unique:     index.Unique(),
		Background: index.Background(),
		Sparse:     index.Sparse(),
		Key:        index.Fields(),
	}

	return rcv.db.DB(databaseName).C(collectionName).EnsureIndex(i)
}
