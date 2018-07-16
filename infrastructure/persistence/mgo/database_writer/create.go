package database_writer

import (
	"fmt"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2"
)

// Creates a new "DeleteMe" collection to create DB.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
func (rcv *databaseWriter) Create(name value.DBName) error {
	if err := name.Valid(); err != nil {
		return err
	}

	database := rcv.db.DB(string(name))
	collectionName := "DeleteMe"

	collection := mgo.Collection{
		Database: database,
		Name:     collectionName,
		FullName: fmt.Sprintf("%s.%s", name, collectionName),
	}

	return collection.Create(&mgo.CollectionInfo{})
}
