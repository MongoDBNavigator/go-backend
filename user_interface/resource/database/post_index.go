package database

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to post index
func (rcv *databaseResource) postIndex(w http.ResponseWriter, r *http.Request) {
	var dbName value.DBName
	var collName value.CollName

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, nil, nil); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	var postRequest representation.PostIndex

	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	index := model.NewIndex(
		postRequest.Name,
		postRequest.Unique,
		postRequest.Background,
		postRequest.Sparse,
		postRequest.Fields,
		postRequest.PartialFilterExpression,
	)

	if err := rcv.indexWriter.Create(dbName, collName, index); err != nil {
		rcv.writeErrorResponse(w, http.StatusConflict, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
