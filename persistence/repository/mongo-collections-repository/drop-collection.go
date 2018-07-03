package mongo_collections_repository

//
// Removes a collection or view from the database.
// https://docs.mongodb.com/manual/reference/method/db.collection.drop/
//
func (rcv *collectionsRepository) DropCollection(databaseName string, collectionName string) error {
	return rcv.db.DB(databaseName).C(collectionName).DropCollection()
}
