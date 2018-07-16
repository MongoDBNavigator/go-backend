package document_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2"
)

// Fetch documents with pagination and filters
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) ReadeAll(conditions *value.ReadAllDocConditions) ([]interface{}, error) {
	collection := rcv.db.DB(string(conditions.DbName())).C(string(conditions.CollName()))
	var result []interface{}
	var query *mgo.Query

	if len(conditions.Filter()) != 0 {
		query = collection.Find(conditions.Filter())
	} else {
		query = collection.Find(nil)
	}

	query.Limit(conditions.Limit()).Skip(conditions.Skip())

	if len(conditions.Sort()) != 0 {
		query.Sort(conditions.Sort()...)
	}

	if err := query.All(&result); err != nil {
		return nil, err
	}

	return result, nil
}
