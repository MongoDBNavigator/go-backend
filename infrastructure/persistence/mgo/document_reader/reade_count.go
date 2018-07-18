package document_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2"
)

// Get documents count by filters
// https://docs.mongodb.com/manual/reference/method/db.collection.find/
func (rcv *documentReader) ReadCount(conditions *value.ReadAllDocConditions) (int, error) {
	collection := rcv.db.DB(string(conditions.DbName())).C(string(conditions.CollName()))
	var query *mgo.Query

	if len(conditions.Filter()) != 0 {
		query = collection.Find(conditions.Filter())
	} else {
		query = collection.Find(nil)
	}

	return query.Count()
}
