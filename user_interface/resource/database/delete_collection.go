package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to delete collection
func (rcv *databaseResource) deleteCollection(w http.ResponseWriter, r *http.Request) {
	var dbName value.DBName
	var collName value.CollName

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, nil, nil); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := rcv.collectionsWriter.Delete(dbName, collName); err != nil {
		rcv.writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
