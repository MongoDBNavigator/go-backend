package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) getDocument(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName, documentId string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, &documentId, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	record, err := rcv.documentsRepository.GetById(databaseName, collectionName, documentId)

	if err != nil {
		response.WriteHeaderAndEntity(http.StatusNotFound, representation.Error{Message: err.Error()})
		return
	}

	response.WriteEntity(record)
}
