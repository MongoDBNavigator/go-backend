package database_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
)

func (rcv *databaseResource) getCollections(request *restful.Request, response *restful.Response) {
	var databaseName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, nil, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	collections, err := rcv.collectionsRepository.GetCollectionsByDatabase(databaseName)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	response.WriteEntity(transformer.CollectionsToView(collections))
}
