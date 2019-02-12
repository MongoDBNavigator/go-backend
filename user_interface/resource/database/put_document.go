package database

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/request"
)

// Method to put collection
func (rcv *databaseResource) putDocument(w http.ResponseWriter, r *http.Request) {
	var dbName value.DBName
	var collName value.CollName
	var docId value.DocId

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := request.ExtractParametersFromRequest(r, &dbName, &collName, &docId, nil); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if _, err := rcv.documentReader.Read(dbName, collName, docId); err != nil {
		rcv.writeErrorResponse(w, http.StatusNotFound, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := rcv.documentWriter.Update(dbName, collName, docId, body); err != nil {
		rcv.writeErrorResponse(w, http.StatusConflict, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
