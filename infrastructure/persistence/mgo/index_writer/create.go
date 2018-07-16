package index_writer

import (
	"gopkg.in/mgo.v2"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Creates indexes on collections.
// https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/
func (rcv *indexWriter) Create(dbName value.DBName, collName value.CollName, index *model.Index) error {
	return rcv.db.DB(string(dbName)).C(string(collName)).EnsureIndex(mgo.Index{
		Name:       index.Name(),
		Unique:     index.Unique(),
		Background: index.Background(),
		Sparse:     index.Sparse(),
		Key:        index.Fields(),
	})
}
