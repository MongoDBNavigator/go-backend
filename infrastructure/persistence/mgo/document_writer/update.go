package document_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2/bson"
)

// Update document
// https://docs.mongodb.com/manual/tutorial/update-documents/
func (rcv *documentWriter) Update(dbName value.DBName, collName value.CollName, docId value.DocId, doc interface{}) error {
	var id interface{}

	if bson.IsObjectIdHex(string(docId)) {
		id = bson.ObjectIdHex(string(docId))
	} else {
		id = docId
	}

	return rcv.db.DB(string(dbName)).C(string(collName)).UpdateId(id, doc)
}
