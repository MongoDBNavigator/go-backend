package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/response"
	"github.com/emicklei/go-restful"
)

// Method to get documents list in json
func (rcv *databaseResource) getDocuments(req *restful.Request, res *restful.Response) {
	listConditions, err := transformer.RequestToListConditions(req)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	records, err := rcv.documentReader.ReadAll(listConditions)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	total, err := rcv.documentReader.ReadCount(listConditions)

	if err != nil {
		res.WriteHeaderAndEntity(http.StatusInternalServerError, representation.Error{Message: err.Error()})
		return
	}

	res.WriteEntity(response.DocumentsToView(records, total))
}
