package document_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2/bson"
)

// Fetch document by ID
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) Read(dbName value.DBName, collName value.CollName, docId value.DocId) (interface{}, error) {
	var result interface{}
	var id interface{}

	if bson.IsObjectIdHex(string(docId)) {
		id = bson.ObjectIdHex(string(docId))
	} else {
		id = docId
	}

	if err := rcv.db.DB(string(dbName)).C(string(collName)).FindId(id).One(&result); err != nil {
		return nil, err
	}

	return result, nil
}
