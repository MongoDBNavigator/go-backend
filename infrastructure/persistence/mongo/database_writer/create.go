package database_writer

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Creates a new "DeleteMe" collection to create DB.
// https://docs.mongodb.com/manual/reference/method/db.createCollection/
func (rcv *databaseWriter) Create(name value.DBName) error {
	res := rcv.db.Database(string(name)).RunCommand(context.Background(), bson.D{{"create", "DeleteMe"}})

	if res.Err() != nil {
		log.Println(res.Err())
	}

	return res.Err()
}
