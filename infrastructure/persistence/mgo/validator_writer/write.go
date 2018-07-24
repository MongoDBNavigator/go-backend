package validator_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2/bson"
)

func (rcv *validatorWriter) Write(dbName value.DBName, collName value.CollName) error {
	var res bson.M
	rcv.db.Run(bson.M{
		"collMod": &collName,
		"validator": bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"phone", "name"},
				"properties": bson.M{
					"phone": bson.M{
						"bsonType":    "string",
						"description": "Phone must be a string and is required",
					},
					"name": bson.M{
						"bsonType":    "string",
						"description": "Name must be a string and is required",
					},
				},
			},
		},
	}, &res)

	return nil
}
