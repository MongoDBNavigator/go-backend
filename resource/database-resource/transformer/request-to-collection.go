package transformer

import (
	"errors"
	"strings"

	"github.com/emicklei/go-restful"
)

func RequestToCollectionName(request *restful.Request) (string, error) {
	collectionName := request.PathParameter("collectionName")

	if len(strings.TrimSpace(collectionName)) == 0 {
		return "", errors.New("Collection name should not be blank.")
	}

	return collectionName, nil
}
