package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/response"
	"github.com/emicklei/go-restful"
)

// Method to get databases list in json
func (rcv *databaseResource) getDatabases(req *restful.Request, res *restful.Response) {
	databases, err := rcv.databaseReader.ReadAll()

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	res.WriteEntity(response.DatabasesToView(databases))
}
