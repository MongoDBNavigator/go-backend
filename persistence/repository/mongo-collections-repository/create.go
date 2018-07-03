package mongo_collections_repository

import (
	"fmt"

	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"gopkg.in/mgo.v2"
)

//
// Creates a new collection or view.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
//
func (rcv *collectionsRepository) Create(collectionInfo *repository.CollectionInfo) error {
	database := rcv.db.DB(collectionInfo.DatabaseName)

	collection := mgo.Collection{
		Database: database,
		Name:     collectionInfo.Name,
		FullName: fmt.Sprintf("%s.%s", collectionInfo.DatabaseName, collectionInfo.Name),
	}

	return collection.Create(&mgo.CollectionInfo{
		DisableIdIndex: collectionInfo.DisableIdIndex,
		ForceIdIndex:   collectionInfo.ForceIdIndex,
		MaxBytes:       collectionInfo.MaxBytes,
		MaxDocs:        collectionInfo.MaxDocs,
		Capped:         collectionInfo.Capped,
	})
}
