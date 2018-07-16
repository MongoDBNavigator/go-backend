package system_info_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/system/repository"
	"gopkg.in/mgo.v2"
)

type systemInfoReader struct {
	db  *mgo.Session
	url string
}

// Constructor for systemInfoReader
func New(db *mgo.Session, url string) repository.SystemInfoReader {
	return &systemInfoReader{
		db:  db,
		url: url,
	}
}
