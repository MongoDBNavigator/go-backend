package transformer

import (
	"errors"
	"strings"

	"fmt"
	"net/url"

	"github.com/emicklei/go-restful"
)

func RequestToSort(request *restful.Request) ([]string, error) {
	params, err := url.ParseQuery(request.Request.URL.RawQuery)

	if err != nil {
		return nil, err
	}

	var sorts []string

	if rawSorts, ok := params["sort[]"]; ok {
		sorts = make([]string, len(rawSorts))

		for i, sort := range rawSorts {
			if len(strings.TrimSpace(sort)) == 0 {
				return nil, errors.New("sort[] value should not be blank.")
			}

			if strings.HasPrefix(sort, "-") == false {
				sort = fmt.Sprintf("+%s", sort)
			}

			sorts[i] = sort
		}
	}

	return sorts, nil
}
