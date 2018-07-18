package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/transformer/request"
	"github.com/emicklei/go-restful"
)

// Method to delete index
func (rcv *databaseResource) deleteIndex(req *restful.Request, res *restful.Response) {
	var dbName value.DBName
	var collName value.CollName
	var indexName value.IndexName

	if err := request.ExtractParametersFromRequest(req, &dbName, &collName, nil, &indexName); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.indexWriter.Delete(dbName, collName, indexName); err != nil {
		res.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusAccepted)
}
