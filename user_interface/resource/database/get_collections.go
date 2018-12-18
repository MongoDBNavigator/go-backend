package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/response"
	"github.com/emicklei/go-restful"
)

// Method to get collections list in json
func (rcv *databaseResource) getCollections(req *restful.Request, res *restful.Response) {
	dbName, err := request.ExtractDatabaseName(req)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	collections, err := rcv.collectionsReader.ReadAll(dbName)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	res.WriteEntity(response.CollectionsToView(collections))
}
