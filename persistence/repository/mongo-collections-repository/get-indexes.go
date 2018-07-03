package mongo_collections_repository

//
// Returns an array of documents that describe the existing indexes on a collection.
// https://docs.mongodb.com/manual/reference/method/db.collection.getIndexes/#db.collection.getIndexes
//
func (rcv *collectionsRepository) GetIndexes(databaseName string, collectionName string) ([]string, error) {

	return []string{}, nil
}
