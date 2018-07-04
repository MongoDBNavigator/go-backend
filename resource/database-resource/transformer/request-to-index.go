package transformer

import (
	"errors"
	"strings"

	"github.com/emicklei/go-restful"
)

func RequestToIndex(request *restful.Request) (string, error) {
	recordId := request.PathParameter("indexName")

	if len(strings.TrimSpace(recordId)) == 0 {
		return "", errors.New("Index name should not be blank.")
	}

	return recordId, nil
}
