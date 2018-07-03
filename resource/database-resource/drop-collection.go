package database_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
)

func (rcv *databaseResource) dropCollection(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	if err := rcv.collectionsRepository.DropCollection(databaseName, collectionName); err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusAccepted)
}
