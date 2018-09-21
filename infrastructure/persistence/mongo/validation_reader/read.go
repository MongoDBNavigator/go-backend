package validation_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

func (rcv *validationReader) Read(dbName value.DBName, collName value.CollName) (*model.Validation, error) {
	return nil, nil
}
