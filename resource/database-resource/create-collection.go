package database_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
)

func (rcv *databaseResource) createCollection(request *restful.Request, response *restful.Response) {
	databaseName, err := transformer.RequestToDatabaseName(request)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	postRequest := new(representation.PostCollection)

	if err := request.ReadEntity(&postRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	collectionInfo := &repository.CollectionInfo{
		DatabaseName: databaseName,
		Name:         postRequest.Name,
	}

	if err := rcv.collectionsRepository.Create(collectionInfo); err != nil {
		response.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusCreated)
}
