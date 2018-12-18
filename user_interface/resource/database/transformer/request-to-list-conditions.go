package transformer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/emicklei/go-restful"
)

func RequestToListConditions(req *restful.Request) (*value.ReadAllDocConditions, error) {
	var dbName value.DBName
	var collName value.CollName

	if err := request.ExtractParametersFromRequest(req, &dbName, &collName, nil, nil); err != nil {
		return nil, err
	}

	limit, err := request.ExtractLimit(req)

	if err != nil {
		return nil, err
	}

	skip, err := request.ExtractSkip(req)

	if err != nil {
		return nil, err
	}

	sorts, err := request.ExtractSort(req)

	if err != nil {
		return nil, err
	}

	filters, err := request.ExtractFilter(req)

	if err != nil {
		return nil, err
	}

	return value.NewReadAllDocConditions(dbName, collName, limit, skip, sorts, filters), nil
}
