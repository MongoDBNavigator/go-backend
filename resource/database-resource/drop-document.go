package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/resource/database-resource/representation"
	"github.com/MongoDBNavigator/go-backend/resource/database-resource/transformer"
	"github.com/emicklei/go-restful"
)

func (rcv *databaseResource) dropDocument(request *restful.Request, response *restful.Response) {
	var databaseName, collectionName, documentId string

	if err := transformer.ExtractParametersFromRequest(request, &databaseName, &collectionName, &documentId, nil); err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, err)
		return
	}

	if err := rcv.documentsRepository.Drop(databaseName, collectionName, documentId); err != nil {
		response.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	response.WriteHeader(http.StatusAccepted)
}
