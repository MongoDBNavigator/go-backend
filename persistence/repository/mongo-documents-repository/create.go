package mongo_documents_repository

func (rcv *documentsRepository) Create(databaseName string, collectionName string, record interface{}) error {
	return rcv.db.DB(databaseName).C(collectionName).Insert(record)
}
