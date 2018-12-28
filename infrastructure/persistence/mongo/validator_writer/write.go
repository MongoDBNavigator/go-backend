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

	_, err = rcv.db.Database(string(dbName)).RunCommand(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("collMod", string(collName)),
			bson.EC.SubDocument("validator", jsonSchema),
			bson.EC.String("validationLevel", string(validation.ValidationLevel())),
			bson.EC.String("validationAction", string(validation.ValidationAction())),
		),
	)

	if err != nil {
		log.Println(err)
	}

	return err
}
