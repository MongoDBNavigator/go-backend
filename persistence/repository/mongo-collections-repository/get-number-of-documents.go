package mongo_collections_repository

//
// Returns the count of documents that would match a find() query for the collection or view.
// https://docs.mongodb.com/manual/reference/method/db.collection.count/#db.collection.count
//
func (rcv *collectionsRepository) GetNumberOfDocuments(databaseName string, collectionName string) (int64, error) {
	count, err := rcv.db.DB(databaseName).C(collectionName).Count()

	if err != nil {
		return 0, nil
	}

	return int64(count), nil
}
