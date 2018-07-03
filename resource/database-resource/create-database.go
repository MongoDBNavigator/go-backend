package database_resource

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/MongoDBNavigator/go-backend/persistence/repository"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
)

func (rcv *databaseResource) createDatabase(request *restful.Request, response *restful.Response) {
	postRequest := new(representation.PostDatabase)

	if err := request.ReadEntity(&postRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	collectionInfo := &repository.CollectionInfo{
		DatabaseName: postRequest.Name,
		Name:         "DeleteMe",
	}

	if err := rcv.collectionsRepository.Create(collectionInfo); err != nil {
		response.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusCreated)
}
