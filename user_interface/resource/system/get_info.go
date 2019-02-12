package system

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system/representation"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/system/transformer"
)

// Get system info (version, CPU architecture, etc.)
func (rcv *systemResource) getInfo(w http.ResponseWriter, r *http.Request) {
	info, err := rcv.systemInfoReader.Reade()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(representation.Error{Message: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transformer.InfoToView(info)); err != nil {
		log.Println(err)
	}
}
