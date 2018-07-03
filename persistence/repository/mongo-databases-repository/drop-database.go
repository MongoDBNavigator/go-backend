package mongo_databases_repository

//
// Removes the current database, deleting the associated data files.
// https://docs.mongodb.com/manual/reference/method/db.dropDatabase/
//
func (rcv *databasesRepository) DropDatabase(name string) error {
	return rcv.db.DB(name).DropDatabase()
}
