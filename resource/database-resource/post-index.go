package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/persistence/model"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) postIndex(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, nil, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	postRequest := new(representation.PostIndex)

	if err := request.ReadEntity(&postRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	index := model.NewIndex(
		postRequest.Name,
		postRequest.Unique,
		postRequest.Background,
		postRequest.Sparse,
		postRequest.Fields,
	)

	if err := rcv.collectionsRepository.CreateIndex(databaseName, collectionName, index); err != nil {
		response.WriteHeaderAndEntity(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
