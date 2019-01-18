package request

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
)

// Extract skip parameter from query
func ExtractSkip(request *restful.Request) (int, error) {
	if len(strings.TrimSpace(request.QueryParameter("skip"))) == 0 {
		return 0, nil
	}

	skip, err := strconv.Atoi(request.QueryParameter("skip"))

	if err != nil {
		return 0, errors.New("skip parameter should be of type integer")
	}

	if skip < 0 {
		return 0, errors.New("skip parameter should be 0 or more")
	}

	return skip, nil
}

// Extract limit parameter from query
func ExtractLimit(request *restful.Request) (int, error) {
	if len(strings.TrimSpace(request.QueryParameter("limit"))) == 0 {
		return 10, nil
	}

	limit, err := strconv.Atoi(request.QueryParameter("limit"))

	if err != nil {
		return 0, errors.New("limit parameter should be of type integer")
	}

	if limit > 1000 {
		return 0, errors.New("limit parameter should be 1000 or less")
	}

	if limit < 1 {
		return 0, errors.New("limit parameter should be 1 or more")
	}

	return limit, nil
}

// Extract sorts parameter from query
func ExtractSort(request *restful.Request) (map[string]int, error) {
	params, err := url.ParseQuery(request.Request.URL.RawQuery)

	if err != nil {
		return nil, err
	}

	sorts := make(map[string]int)

	if rawSorts, ok := params["sort[]"]; ok {
		for _, sort := range rawSorts {
			sort = strings.TrimSpace(sort)
			if len(sort) == 0 {
				return nil, errors.New("sort[] value should not be blank")
			}

			var direction int

			if strings.HasPrefix(sort, "-") == true {
				direction = -1
				sort = strings.TrimPrefix(sort, "-")
			} else {
				direction = 1
				sort = strings.TrimPrefix(sort, "+")
			}

			sorts[sort] = direction
		}
	}

	return sorts, nil
}

// Extract filters parameter from query
func ExtractFilter(request *restful.Request) ([]byte, error) {
	if len(strings.TrimSpace(request.QueryParameter("filter"))) == 0 {
		return nil, nil
	}

	return []byte(request.QueryParameter("filter")), nil
}
