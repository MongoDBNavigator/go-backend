package database_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
)

func (rcv *databaseResource) dropDatabase(request *restful.Request, response *restful.Response) {
	var databaseName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, nil, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	if err := rcv.databasesRepository.DropDatabase(databaseName); err != nil {
		response.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusAccepted)
}
