package validator_writer

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Implementation ValidationWriter.Write()
func (rcv *validatorWriter) Write(dbName value.DBName, collName value.CollName, validation *model.Validation) error {
	jsonSchema, err := buildJsonSchema(validation)

	if err != nil {
		log.Println(err)
		return err
	}

	collModResult := rcv.db.Database(string(dbName)).RunCommand(
		context.Background(),
		bson.D{
			{"collMod", string(collName)},
			{"validator", jsonSchema},
			{"validationLevel", string(validation.ValidationLevel())},
			{"validationAction", string(validation.ValidationAction())},
		},
	)

	if collModResult.Err() != nil {
		log.Println(collModResult.Err())
	}

	return collModResult.Err()
}
