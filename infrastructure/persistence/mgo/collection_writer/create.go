package collection_writer

import (
	"fmt"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2"
)

// Creates a new collection or view.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
func (rcv *collectionWriter) Create(dbName value.DBName, collName value.CollName) error {
	if err := collName.Valid(); err != nil {
		return err
	}

	database := rcv.db.DB(string(dbName))

	collection := mgo.Collection{
		Database: database,
		Name:     string(collName),
		FullName: fmt.Sprintf("%s.%s", dbName, collName),
	}

	return collection.Create(&mgo.CollectionInfo{})
}
