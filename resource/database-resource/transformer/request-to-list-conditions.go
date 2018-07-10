package transformer

import (
	"errors"
	"strconv"
	"strings"

	"net/url"

	"fmt"

	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"github.com/emicklei/go-restful"
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

	params, err := url.ParseQuery(request.Request.URL.RawQuery)

	if err != nil {
		return nil, err
	}

	//fmt.Println(params["filter[]"])

	var sorts []string

	if _, ok := params["sort[]"]; ok {
		sorts = make([]string, len(params["sort[]"]))

		for i, sort := range params["sort[]"] {
			if len(strings.TrimSpace(sort)) != 0 {
				return nil, errors.New("sort[] value should not be blank.")
			}

			if strings.HasPrefix(sort, "-") == false {
				sort = fmt.Sprintf("+%s", sort)
			}

			sorts[i] = sort
		}
	}

	return repository.NewGetListConditions(databaseName, collectionName, limit, skip, sorts), nil
}
