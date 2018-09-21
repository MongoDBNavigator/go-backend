package collection_writer

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Removes a collection or view from the database.
// https://docs.mongodb.com/manual/reference/method/db.collection.drop/
func (rcv *collectionWriter) Delete(dbName value.DBName, collName value.CollName) error {
	return nil
}
