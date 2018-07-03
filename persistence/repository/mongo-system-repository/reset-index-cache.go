package mongo_system_repository

//
// Removes all cached query plans for a collection.
// https://docs.mongodb.com/manual/reference/method/PlanCache.clear/
//
func (rcv *systemRepository) ResetIndexCache() {
	rcv.db.ResetIndexCache()
}
