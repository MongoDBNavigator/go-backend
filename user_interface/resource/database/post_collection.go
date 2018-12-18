package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/emicklei/go-restful"
)

// Method to post collection
func (rcv *databaseResource) postCollection(req *restful.Request, res *restful.Response) {
	dbName, err := request.ExtractDatabaseName(req)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	postRequest := new(representation.PostCollection)

	if err := req.ReadEntity(&postRequest); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.collectionsWriter.Create(dbName, postRequest.Name); err != nil {
		res.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusCreated)
}
