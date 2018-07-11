package transformer

import (
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"github.com/emicklei/go-restful"
)

func RequestToListConditions(request *restful.Request) (*repository.GetListConditions, error) {
	var databaseName, collectionName string

	if err := ExtractParametersFromRequest(request, &databaseName, &collectionName, nil, nil); err != nil {
		return nil, err
	}

	limit, err := RequestToLimit(request)

	if err != nil {
		return nil, err
	}

	skip, err := RequestToSkip(request)

	if err != nil {
		return nil, err
	}

	sorts, err := RequestToSort(request)

	if err != nil {
		return nil, err
	}

	filters, err := RequestToFilter(request)

	if err != nil {
		return nil, err
	}

	return repository.NewGetListConditions(databaseName, collectionName, limit, skip, sorts, filters), nil
}
