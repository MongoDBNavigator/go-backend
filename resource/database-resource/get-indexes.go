package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) getIndexes(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, nil, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	indexes, err := rcv.collectionsRepository.GetIndexes(databaseName, collectionName)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	response.WriteEntity(transformer.IndexesToView(indexes))
}
