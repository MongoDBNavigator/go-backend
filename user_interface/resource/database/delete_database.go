package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/emicklei/go-restful"
)

// Method to delete database
func (rcv *databaseResource) deleteDatabase(req *restful.Request, res *restful.Response) {
	dbName, err := request.ExtractDatabaseName(req)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.databaseWriter.Delete(dbName); err != nil {
		res.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusAccepted)
}
