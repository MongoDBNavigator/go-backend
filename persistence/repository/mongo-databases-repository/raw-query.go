package mongo_databases_repository

func (rcv *databasesRepository) RawQuery(jsonQuery string) (interface{}, error) {
	var result interface{}

	if err := rcv.db.Run(jsonQuery, result); err != nil {
		return nil, err
	}

	return result, nil
}
