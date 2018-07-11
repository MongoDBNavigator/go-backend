package transformer

import (
	"errors"
	"strings"

	"strconv"

	"github.com/emicklei/go-restful"
)

func RequestToSkip(request *restful.Request) (int, error) {
	if len(strings.TrimSpace(request.QueryParameter("skip"))) == 0 {
		return 0, errors.New("Skip parameter should not be blank.")
	}

	skip, err := strconv.Atoi(request.QueryParameter("skip"))

	if err != nil {
		return 0, errors.New("Skip parameter should be of type integer.")
	}

	if skip < 0 {
		return 0, errors.New("Skip parameter should be 0 or more.")
	}

	return skip, nil
}
