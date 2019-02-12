package database

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to post document
func (rcv *databaseResource) postDocument(w http.ResponseWriter, r *http.Request) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := rcv.documentWriter.Create(dbName, collName, body); err != nil {
		rcv.writeErrorResponse(w, http.StatusConflict, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
