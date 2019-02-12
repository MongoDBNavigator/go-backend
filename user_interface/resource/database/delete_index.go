package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to delete index
func (rcv *databaseResource) deleteIndex(w http.ResponseWriter, r *http.Request) {
	var dbName value.DBName
	var collName value.CollName
	var indexName value.IndexName

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, nil, &indexName); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := rcv.indexWriter.Delete(dbName, collName, indexName); err != nil {
		rcv.writeErrorResponse(w, http.StatusConflict, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
