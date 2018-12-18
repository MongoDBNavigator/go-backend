package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/emicklei/go-restful"
)

// Method to post document
func (rcv *databaseResource) postDocument(req *restful.Request, res *restful.Response) {
	var dbName value.DBName
	var collName value.CollName

	if err := request.ExtractParametersFromRequest(req, &dbName, &collName, nil, nil); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	postRequest := new(interface{})

	if err := req.ReadEntity(&postRequest); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.documentWriter.Create(dbName, collName, postRequest); err != nil {
		res.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusCreated)
}
