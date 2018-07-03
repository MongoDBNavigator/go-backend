package transformer

import (
	"errors"
	"strings"

	"github.com/emicklei/go-restful"
)

func RequestToDatabaseName(request *restful.Request) (string, error) {
	databaseName := request.PathParameter("databaseName")

	if len(strings.TrimSpace(databaseName)) == 0 {
		return "", errors.New("Database name should not be blank.")
	}

	return databaseName, nil
}
