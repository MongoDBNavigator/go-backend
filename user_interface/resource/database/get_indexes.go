package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/response"
)

// Method to get indexes list in json
func (rcv *databaseResource) getIndexes(w http.ResponseWriter, r *http.Request) {
	var dbName value.DBName
	var collName value.CollName

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, nil, nil); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	indexes, err := rcv.indexReader.ReadAll(dbName, collName)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	rcv.writeResponse(w, http.StatusOK, response.IndexesToView(indexes))
}
