package database

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
)

// Method to post database
func (rcv *databaseResource) postDatabase(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var postRequest representation.PostDatabase

	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := rcv.databaseWriter.Create(postRequest.Name); err != nil {
		rcv.writeErrorResponse(w, http.StatusConflict, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
