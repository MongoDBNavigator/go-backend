package transformer

import (
	"errors"
	"strings"

	"strconv"

	"github.com/emicklei/go-restful"
)

func RequestToLimit(request *restful.Request) (int, error) {
	if len(strings.TrimSpace(request.QueryParameter("limit"))) == 0 {
		return 0, errors.New("Limit parameter should not be blank.")
	}

	limit, err := strconv.Atoi(request.QueryParameter("limit"))

	if err != nil {
		return 0, errors.New("Limit parameter should be of type integer.")
	}

	if limit > 1000 {
		return 0, errors.New("Limit parameter should be 1000 or less.")
	}

	if limit < 1 {
		return 0, errors.New("Limit parameter should be 1 or more.")
	}

	return limit, nil
}
