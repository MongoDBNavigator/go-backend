package mongo_documents_repository

import (
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"gopkg.in/mgo.v2"
)

//
// Returns the count of documents that would match a find() query for the collection or view.
// https://docs.mongodb.com/manual/reference/method/db.collection.count/#db.collection.count
//
func (rcv *documentsRepository) GetNumberOfDocuments(conditions *repository.GetListConditions) (int, error) {
	collection := rcv.db.DB(conditions.DatabaseName()).C(conditions.CollectionName())
	var query *mgo.Query

	if len(conditions.Filter()) != 0 {
		query = collection.Find(conditions.Filter())
	} else {
		query = collection.Find(nil)
	}

	return query.Count()
}
