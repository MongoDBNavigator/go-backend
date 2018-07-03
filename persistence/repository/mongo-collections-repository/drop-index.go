package mongo_collections_repository

//
// Drops or removes the specified index from a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.dropIndex/#db.collection.dropIndex
//
func (rcv *collectionsRepository) DropIndex(databaseName string, collectionName string, key ...string) error {
	return rcv.db.DB(databaseName).C(collectionName).DropIndex(key...)
}
