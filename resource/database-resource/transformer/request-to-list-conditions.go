package transformer

import (
	"errors"
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
)

func RequestToListConditions(request *restful.Request) (*repository.GetListConditions, error) {
	databaseName := request.PathParameter("databaseName")

	if len(strings.TrimSpace(databaseName)) == 0 {
		return nil, errors.New("Database name should not be blank.")
	}

	collectionName := request.PathParameter("collectionName")

	if len(strings.TrimSpace(collectionName)) == 0 {
		return nil, errors.New("Collection name should not be blank.")
	}

	if len(strings.TrimSpace(request.QueryParameter("limit"))) == 0 {
		return nil, errors.New("Limit parameter should not be blank.")
	}

	if len(strings.TrimSpace(request.QueryParameter("skip"))) == 0 {
		return nil, errors.New("Skip parameter should not be blank.")
	}

	limit, err := strconv.Atoi(request.QueryParameter("limit"))

	if err != nil {
		return nil, errors.New("Limit parameter should be of type integer.")
	}

	if limit > 1000 {
		return nil, errors.New("Limit parameter should be 1000 or less.")
	}

	if limit < 1 {
		return nil, errors.New("Limit parameter should be 1 or more.")
	}

	skip, err := strconv.Atoi(request.QueryParameter("skip"))

	if err != nil {
		return nil, errors.New("Skip parameter should be of type integer.")
	}

	if skip < 0 {
		return nil, errors.New("Skip parameter should be 0 or more.")
	}

	var sortField, sortDirection string

	sortFieldParam := request.QueryParameter("sortField")
	sortDirectionParam := request.QueryParameter("sortDirection")

	if len(strings.TrimSpace(sortFieldParam)) != 0 && len(strings.TrimSpace(sortDirectionParam)) != 0 {
		if sortDirectionParam != "asc" && sortDirectionParam != "desc" {
			return nil, errors.New("Sort direction should be equal to 'asc' or 'desc'.")
		}

		sortField = sortFieldParam
		sortDirection = sortDirectionParam
	}

	return repository.NewGetListConditions(databaseName, collectionName, limit, skip, sortField, sortDirection), nil
}
