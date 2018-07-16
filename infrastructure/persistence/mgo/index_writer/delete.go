package index_writer

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Drops or removes the specified index from a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.dropIndex/#db.collection.dropIndex
func (rcv *indexWriter) Delete(dbName value.DBName, collName value.CollName, indexName value.IndexName) error {
	return rcv.db.DB(string(dbName)).C(string(collName)).DropIndexName(string(indexName))
}
