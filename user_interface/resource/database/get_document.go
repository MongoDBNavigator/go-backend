package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to get document in json
func (rcv *databaseResource) getDocument(w http.ResponseWriter, r *http.Request) {
	var dbName value.DBName
	var collName value.CollName
	var docId value.DocId

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, &docId, nil); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	document, err := rcv.documentReader.Read(dbName, collName, docId)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusNotFound, err)
		return
	}

	rcv.writeResponse(w, http.StatusOK, document)
}
