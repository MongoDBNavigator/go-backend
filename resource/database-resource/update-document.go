package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) updateDocument(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName, documentId string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, &documentId, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	if _, err := rcv.documentsRepository.GetById(databaseName, collectionName, documentId); err != nil {
		response.WriteHeaderAndEntity(http.StatusNotFound, representation.Error{Message: err.Error()})
		return
	}

	putRequest := new(interface{})

	if err := request.ReadEntity(&putRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.documentsRepository.Update(databaseName, collectionName, documentId, putRequest); err != nil {
		response.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusAccepted)
}
