package validation_reader

import (
	localErr "github.com/MongoDBNavigator/go-backend/domain/database/err"
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
		return nil, localErr.ValidationNotFound
	}

	options := data["cursor"].(bson.M)["firstBatch"].([]interface{})[0].(bson.M)["options"].(bson.M)

	if _, ok := options["validator"]; !ok {
		return nil, localErr.ValidationNotFound
	}

	validator := options["validator"].(bson.M)
	properties := make([]*model.ValidationProperty, 0)

	if jsonSchema, ok := validator["$jsonSchema"]; ok {
		requiredFields := make([]interface{}, 0)

		if reqFields, ok := jsonSchema.(bson.M)["required"]; ok {
			requiredFields = reqFields.([]interface{})
		}

		if props, ok := jsonSchema.(bson.M)["properties"]; ok {
			for field, constraints := range props.(bson.M) {
				var (
					required         bool
					bsonType         string
					enum             []interface{}
					description      string
					minimum          int
					maximum          int
					pattern          string
					maxLength        int
					minLength        int
					exclusiveMaximum bool
					exclusiveMinimum bool
					uniqueItems      bool
					minItems         int
					maxItems         int
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

				if val, ok := constraints.(bson.M)["uniqueItems"]; ok {
					uniqueItems = val.(bool)
				}

				if val, ok := constraints.(bson.M)["enum"]; ok {
					enum = val.([]interface{})
				}

				if val, ok := constraints.(bson.M)["exclusiveMaximum"]; ok {
					exclusiveMaximum = val.(bool)
				}

				if val, ok := constraints.(bson.M)["exclusiveMinimum"]; ok {
					exclusiveMinimum = val.(bool)
				}

				for _, v := range requiredFields {
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
					minItems,
					maxItems,
					exclusiveMaximum,
					exclusiveMinimum,
					uniqueItems,
				))
			}
		}
	}

	validationLevel := value.ValidationLevel(options["validationLevel"].(string))
	validationAction := value.ValidationAction(options["validationAction"].(string))

	return model.NewValidation(validationLevel, validationAction, properties), nil
}
