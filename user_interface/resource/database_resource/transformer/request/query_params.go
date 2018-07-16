package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	restful "github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2/bson"
)

func ExtractSkip(request *restful.Request) (int, error) {
	if len(strings.TrimSpace(request.QueryParameter("skip"))) == 0 {
		return 0, errors.New("skip parameter should not be blank")
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

func ExtractLimit(request *restful.Request) (int, error) {
	if len(strings.TrimSpace(request.QueryParameter("limit"))) == 0 {
		return 0, errors.New("limit parameter should not be blank")
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

func ExtractSort(request *restful.Request) ([]string, error) {
	params, err := url.ParseQuery(request.Request.URL.RawQuery)

	if err != nil {
		return nil, err
	}

	var sorts []string

	if rawSorts, ok := params["sort[]"]; ok {
		sorts = make([]string, len(rawSorts))

		for i, sort := range rawSorts {
			sort = strings.TrimSpace(sort)
			if len(sort) == 0 {
				return nil, errors.New("sort[] value should not be blank")
			}

			if strings.HasPrefix(sort, "-") == false {
				sort = fmt.Sprintf("+%s", sort)
			}

			sorts[i] = sort
		}
	}

	return sorts, nil
}

func ExtractFilter(request *restful.Request) (bson.M, error) {
	params, err := url.ParseQuery(request.Request.URL.RawQuery)

	if err != nil {
		return nil, err
	}

	filters := make(bson.M)

	if rawFilters, ok := params["filter[]"]; ok {
		for _, filterString := range rawFilters {
			if len(strings.TrimSpace(filterString)) == 0 {
				return nil, errors.New("filter[] value should not be blank")
			}

			filter, err := parseFilter(filterString)

			if err != nil {
				return nil, err
			}

			for k, v := range filter {
				filters[k] = v
			}
		}
	}

	return filters, nil
}

func parseFilter(data string) (bson.M, error) {
	if !strings.Contains(data, ":") {
		return nil, errors.New("filter[] has bad formant (field:value)")
	}

	dataSlice := strings.SplitN(data, ":", 2)

	if len(strings.TrimSpace(dataSlice[0])) == 0 {
		return nil, errors.New("filter kay should not be blank")
	}

	if len(strings.TrimSpace(dataSlice[1])) == 0 {
		return nil, errors.New("filter value should not be blank")
	}

	jsonFormat := `{"%s":%s}`

	if reflect.TypeOf(dataSlice[1]).String() == "string" {
		jsonFormat = `{"%s":"%s"}`
	}

	jsonData := fmt.Sprintf(jsonFormat, dataSlice[0], dataSlice[1])

	var filter map[string]interface{}

	if err := json.Unmarshal([]byte(jsonData), &filter); err != nil {
		return nil, err
	}

	return convertToBson(filter), nil
}

func convertToBson(data map[string]interface{}) bson.M {
	filter := make(bson.M)

	if regex, ok := data["$regex"]; ok {
		var options string
		if opts, ok := data["$options"]; ok {
			options = opts.(string)
		}

		filter["$regex"] = bson.RegEx{
			Pattern: regex.(string),
			Options: options,
		}

		return filter
	}

	for k, v := range data {
		vType := reflect.TypeOf(v)

		if vType.String() == "map[string]interface {}" {
			filter[k] = convertToBson(v.(map[string]interface{}))
		} else if vType.String() == "[]interface {}" {
			vSlice := v.([]interface{})
			filter[k] = make([]interface{}, len(vSlice))
			for i, j := range vSlice {
				filter[k].([]interface{})[i] = convertToBson(j.(map[string]interface{}))
			}
		} else {
			filter[k] = v
		}
	}

	return filter
}
