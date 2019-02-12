package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/response"
)

// Method to get databases list in json
func (rcv *databaseResource) getDatabases(w http.ResponseWriter, r *http.Request) {
	databases, err := rcv.databaseReader.ReadAll()

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	rcv.writeResponse(w, http.StatusOK, response.DatabasesToView(databases))
}
