package mongo_documents_repository

import (
	"fmt"

	"github.com/MongoDBNavigator/go-backend/persistence/repository"
)

func (rcv *documentsRepository) GetAll(conditions *repository.GetListConditions) ([]interface{}, error) {
	var result []interface{}
	query := rcv.db.DB(conditions.GetDatabaseName()).
		C(conditions.GetCollectionName()).
		Find(nil).
		Limit(conditions.GetLimit()).
		Skip(conditions.GetSkip())

	fmt.Println(conditions.GetSort())

	if len(conditions.GetSort()) != 0 {
		query.Sort(conditions.GetSort()...)
	}

	if err := query.All(&result); err != nil {
		return nil, err
	}

	return result, nil
}
