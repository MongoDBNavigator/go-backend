package mongo_documents_repository

import "gopkg.in/mgo.v2/bson"

func (rcv *documentsRepository) GetById(databaseName string, collectionName string, recordId string) (interface{}, error) {
	var result interface{}
	var id interface{}

	if bson.IsObjectIdHex(recordId) {
		id = bson.ObjectIdHex(recordId)
	} else {
		id = recordId
	}

	query := rcv.db.DB(databaseName).C(collectionName).FindId(id)

	if err := query.One(&result); err != nil {
		return nil, err
	}

	return result, nil
}
