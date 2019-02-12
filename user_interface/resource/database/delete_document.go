package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to delete document
func (rcv *databaseResource) deleteDocument(w http.ResponseWriter, r *http.Request) {
	var dbName value.DBName
	var collName value.CollName
	var docId value.DocId

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, &docId, nil); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := rcv.documentWriter.Delete(dbName, collName, docId); err != nil {
		rcv.writeErrorResponse(w, http.StatusConflict, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
