package document_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2/bson"
)

// Drop document
// https://docs.mongodb.com/manual/tutorial/remove-documents/
func (rcv *documentWriter) Delete(dbName value.DBName, collName value.CollName, docId value.DocId) error {
	var id interface{}

	if bson.IsObjectIdHex(string(docId)) {
		id = bson.ObjectIdHex(string(docId))
	} else {
		id = docId
	}

	return rcv.db.DB(string(dbName)).C(string(collName)).RemoveId(id)
}
