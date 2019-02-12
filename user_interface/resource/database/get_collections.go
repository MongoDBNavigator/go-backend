package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/response"
)

// Method to get collections list in json
func (rcv *databaseResource) getCollections(w http.ResponseWriter, r *http.Request) {
	dbName, err := request.ExtractDatabaseName(r)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	collections, err := rcv.collectionsReader.ReadAll(dbName)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	rcv.writeResponse(w, http.StatusOK, response.CollectionsToView(collections))
}
