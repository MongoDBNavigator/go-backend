package database_resource

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/transformer/request"
	"github.com/emicklei/go-restful"
)

// Method to post index
func (rcv *databaseResource) postIndex(req *restful.Request, res *restful.Response) {
	var dbName value.DBName
	var collName value.CollName

	if err := request.ExtractParametersFromRequest(req, &dbName, &collName, nil, nil); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	postRequest := new(representation.PostIndex)

	if err := req.ReadEntity(&postRequest); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	index := model.NewIndex(
		postRequest.Name,
		postRequest.Unique,
		postRequest.Background,
		postRequest.Sparse,
		postRequest.Fields,
	)

	if err := rcv.indexWriter.Create(dbName, collName, index); err != nil {
		res.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusCreated)
}
