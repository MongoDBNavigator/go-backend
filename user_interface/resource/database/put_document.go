package database

import (
	"io/ioutil"
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/emicklei/go-restful"
)

// Method to put collection
func (rcv *databaseResource) putDocument(req *restful.Request, res *restful.Response) {
	var dbName value.DBName
	var collName value.CollName
	var docId value.DocId

	if err := request.ExtractParametersFromRequest(req, &dbName, &collName, &docId, nil); err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if _, err := rcv.documentReader.Read(dbName, collName, docId); err != nil {
		res.WriteHeaderAndEntity(http.StatusNotFound, representation.Error{Message: err.Error()})
		return
	}

	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		res.WriteHeaderAndEntity(http.StatusBadRequest, representation.Error{Message: err.Error()})
		return
	}

	if err := rcv.documentWriter.Update(dbName, collName, docId, body); err != nil {
		res.WriteHeaderAndEntity(http.StatusConflict, representation.Error{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusAccepted)
}
