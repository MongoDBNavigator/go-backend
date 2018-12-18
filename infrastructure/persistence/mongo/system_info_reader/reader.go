package system_info_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/system/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Structure to implement SystemInfoReader interface
type systemInfoReader struct {
	db  *mongo.Client
	url string
}

// Constructor for systemInfoReader
func New(db *mongo.Client, url string) repository.SystemInfoReader {
	return &systemInfoReader{
		db:  db,
		url: url,
	}
}
