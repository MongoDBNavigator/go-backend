package transformer

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

func RequestToListConditions(r *http.Request) (*value.ReadAllDocConditions, error) {
	var dbName value.DBName
	var collName value.CollName

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, nil, nil); err != nil {
		return nil, err
	}

	limit, err := request.ExtractLimit(r)

	if err != nil {
		return nil, err
	}

	skip, err := request.ExtractSkip(r)

	if err != nil {
		return nil, err
	}

	sorts, err := request.ExtractSort(r)

	if err != nil {
		return nil, err
	}

	filters, err := request.ExtractFilter(r)

	if err != nil {
		return nil, err
	}

	return value.NewReadAllDocConditions(dbName, collName, limit, skip, sorts, filters), nil
}
