package mongo_collections_repository

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//
// Creates indexes on collections.
// https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
//
func (rcv *collectionsRepository) CreateIndex(databaseName string, collectionName string) error {

	index := mgo.Index{}

	fmt.Println(index)

	return nil
}
