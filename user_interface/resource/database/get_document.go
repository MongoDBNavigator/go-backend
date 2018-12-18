package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/emicklei/go-restful"
)

// Method to get document in json
func (rcv *databaseResource) getDocument(req *restful.Request, res *restful.Response) {
	var dbName value.DBName
	var collName value.CollName
	var docId value.DocId

	if err := request.ExtractParametersFromRequest(req, &dbName, &collName, &docId, nil); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	document, err := rcv.documentReader.Read(dbName, collName, docId)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusNotFound, representation.Error{Message: err.Error()})
		return
	}

	res.WriteEntity(document)
}
