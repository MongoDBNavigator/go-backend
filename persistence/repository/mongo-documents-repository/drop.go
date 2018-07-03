package mongo_documents_repository

import "gopkg.in/mgo.v2/bson"

func (rcv *documentsRepository) Drop(databaseName string, collectionName string, recordId string) error {
	var id interface{}

	if bson.IsObjectIdHex(recordId) {
		id = bson.ObjectIdHex(recordId)
	} else {
		id = recordId
	}

	return rcv.db.DB(databaseName).C(collectionName).RemoveId(id)
}
