package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to delete database
func (rcv *databaseResource) deleteDatabase(w http.ResponseWriter, r *http.Request) {
	dbName, err := request.ExtractDatabaseName(r)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := rcv.databaseWriter.Delete(dbName); err != nil {
		rcv.writeErrorResponse(w, http.StatusConflict, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
