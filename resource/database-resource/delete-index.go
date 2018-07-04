package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) deleteIndex(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName, indexName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, nil, &indexName); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	if err := rcv.collectionsRepository.DropIndex(databaseName, collectionName, indexName); err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeader(http.StatusAccepted)
}
