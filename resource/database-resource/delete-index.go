package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) deleteIndex(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName, indexName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, nil, &indexName); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.collectionsRepository.DropIndex(databaseName, collectionName, indexName); err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusAccepted)
}
