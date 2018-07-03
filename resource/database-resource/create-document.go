package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) createDocument(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	postRequest := new(interface{})

	if err := request.ReadEntity(&postRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.documentsRepository.Create(databaseName, collectionName, postRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusCreated)
}
