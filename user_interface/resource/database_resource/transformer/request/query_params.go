package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2/bson"
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
				sort = strings.TrimPrefix("sort", "-")
			} else {
				direction = 1
				sort = strings.TrimPrefix("sort", "+")
			}

			sorts[sort] = direction
		}
	}

	return sorts, nil
}

// Extract filters parameter from query
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

// Parse filters parameter
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

	return convertFilterToBson(filter), nil
}

// Convert filter parameter to BSON
func convertFilterToBson(data map[string]interface{}) bson.M {
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
			filter[k] = convertFilterToBson(v.(map[string]interface{}))
		} else if vType.String() == "[]interface {}" {
			vSlice := v.([]interface{})
			filter[k] = make([]interface{}, len(vSlice))
			for i, j := range vSlice {
				filter[k].([]interface{})[i] = convertFilterToBson(j.(map[string]interface{}))
			}
		} else if vType.String() == "string" {
			if bson.IsObjectIdHex(v.(string)) {
				filter[k] = bson.ObjectIdHex(v.(string))
			} else {
				filter[k] = v
			}
		} else {
			filter[k] = v
		}
	}

	return filter
}
