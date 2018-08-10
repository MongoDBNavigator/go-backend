package validator_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2/bson"
)

// Implementation ValidationWriter.Write()
func (rcv *validatorWriter) Write(dbName value.DBName, collName value.CollName, validation *model.Validation) error {
	jsonSchema, err := buildJsonSchema(validation)

	if err != nil {
		return err
	}

	if err := rcv.db.DB(string(dbName)).Run(bson.M{"collMod": &collName, "validator": bson.M{"$jsonSchema": jsonSchema}}, nil); err != nil {
		return err
	}

	return nil
}
