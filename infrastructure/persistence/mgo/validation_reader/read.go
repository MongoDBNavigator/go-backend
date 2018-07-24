package validation_reader

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"gopkg.in/mgo.v2/bson"
)

func (rcv *validationReader) Read(dbName value.DBName, collName value.CollName) (*model.Validation, error) {
	var data bson.M

	if err := rcv.db.DB(string(dbName)).Run(bson.M{"listCollections": 1, "filter": bson.M{"name": collName}}, &data); err != nil {
		return nil, err
	}

	if len(data["cursor"].(bson.M)["firstBatch"].([]interface{})) == 0 {
		return nil, err.ValidationNotFound
	}

	options := data["cursor"].(bson.M)["firstBatch"].([]interface{})[0].(bson.M)["options"].(bson.M)
	validator := options["validator"].(bson.M)
	properties := make([]*model.ValidationProperty, 0)

	if jsonSchema, ok := validator["$jsonSchema"]; ok {
		requiredFields := jsonSchema.(bson.M)["required"]

		for field, constraints := range jsonSchema.(bson.M)["properties"].(bson.M) {
			var (
				required    bool
				bsonType    string
				enum        []interface{}
				description string
				minimum     int
				maximum     int
				pattern     string
				maxLength   int
				minLength   int
			)

			if val, ok := constraints.(bson.M)["bsonType"]; ok {
				bsonType = val.(string)
			}

			if val, ok := constraints.(bson.M)["description"]; ok {
				description = val.(string)
			}

			if val, ok := constraints.(bson.M)["pattern"]; ok {
				pattern = val.(string)
			}

			if val, ok := constraints.(bson.M)["minimum"]; ok {
				minimum = val.(int)
			}

			if val, ok := constraints.(bson.M)["maximum"]; ok {
				maximum = val.(int)
			}

			if val, ok := constraints.(bson.M)["maxLength"]; ok {
				maxLength = val.(int)
			}

			if val, ok := constraints.(bson.M)["minLength"]; ok {
				minLength = val.(int)
			}

			if val, ok := constraints.(bson.M)["enum"]; ok {
				enum = val.([]interface{})
			}

			for _, v := range requiredFields.([]interface{}) {
				if v == field {
					required = true
					break
				}
			}

			properties = append(properties, model.NewValidationProperty(
				field,
				required,
				bsonType,
				enum,
				description,
				minimum,
				maximum,
				pattern,
				maxLength,
				minLength,
			))
		}
	}

	validationLevel := value.ValidationLevel(options["validationLevel"].(string))
	validationAction := value.ValidationAction(options["validationAction"].(string))

	return model.NewValidation(validationLevel, validationAction, properties), nil
}
