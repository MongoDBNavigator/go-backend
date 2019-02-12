package database

import (
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/transformer/response"
)

// Method to get documents list in json
func (rcv *databaseResource) getDocuments(w http.ResponseWriter, r *http.Request) {
	listConditions, err := transformer.RequestToListConditions(r)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	records, err := rcv.documentReader.ReadAll(listConditions)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	total, err := rcv.documentReader.ReadCount(listConditions)

	if err != nil {
		rcv.writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	rcv.writeResponse(w, http.StatusOK, response.DocumentsToView(records, total))
}
