package validator_writer

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Implementation ValidationWriter.Write()
func (rcv *validatorWriter) Write(dbName value.DBName, collName value.CollName, validation *model.Validation) error {
	return nil
}
