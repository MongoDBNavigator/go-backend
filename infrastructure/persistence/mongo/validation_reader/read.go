package validation_reader

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
)

// Read method to read $jsonSchema validation
// JSON Schema is the recommended means of performing schema validation.
// https://docs.mongodb.com/manual/reference/operator/query/jsonSchema/#jsonschema
// http://json-schema.org/
func (rcv *validationReader) Read(dbName value.DBName, collName value.CollName) (*model.Validation, error) {
	listCollectionsResult := rcv.db.Database(string(dbName)).RunCommand(
		context.Background(),
		bson.D{
			{"listCollections", 1},
			{"filter", bson.D{{"name", string(collName)}}},
		},
	)

	if listCollectionsResult.Err() != nil {
		log.Println(listCollectionsResult.Err())
		return nil, listCollectionsResult.Err()
	}

	raw, err := listCollectionsResult.DecodeBytes()

	if err != nil {
		return nil, err
	}

	cursor, err := raw.LookupErr("cursor")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	firstBatch, err := cursor.Document().LookupErr("firstBatch")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	data, err := firstBatch.Array().IndexErr(0)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	options, err := data.Value().Document().LookupErr("options")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	validator, err := options.Document().LookupErr("validator")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	jsonSchema, err := validator.Document().LookupErr("$jsonSchema")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var requiredFields []string

	// collect required fields
	if required, err := jsonSchema.Document().LookupErr("required"); err == nil {
		elements, err := required.Array().Elements()

		if err != nil {
			log.Println(err)
			return nil, err
		}

		requiredFields = make([]string, len(elements))

		for i, e := range elements {
			requiredFields[i] = e.Value().StringValue()
		}
	}

	props, err := jsonSchema.Document().LookupErr("properties")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	rawProps, err := props.Document().Elements()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	properties := make([]*model.ValidationProperty, len(rawProps))

	for i, p := range rawProps {
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

		for _, v := range requiredFields {
			if v == p.Key() {
				required = true
				break
			}
		}

		if element, err := p.Value().Document().LookupErr("bsonType"); err == nil {
			bsonType = element.StringValue()
		}

		if element, err := p.Value().Document().LookupErr("description"); err == nil {
			description = element.StringValue()
		}

		if element, err := p.Value().Document().LookupErr("pattern"); err == nil {
			pattern = element.StringValue()
		}

		if element, err := p.Value().Document().LookupErr("minimum"); err == nil {
			minimum = int(element.Int32())
		}

		if element, err := p.Value().Document().LookupErr("maximum"); err == nil {
			maximum = int(element.Int32())
		}

		if element, err := p.Value().Document().LookupErr("maxLength"); err == nil {
			maxLength = int(element.Int32())
		}

		if element, err := p.Value().Document().LookupErr("minLength"); err == nil {
			minLength = int(element.Int32())
		}

		if element, err := p.Value().Document().LookupErr("uniqueItems"); err == nil {
			uniqueItems = element.Boolean()
		}

		if element, err := p.Value().Document().LookupErr("exclusiveMaximum"); err == nil {
			exclusiveMaximum = element.Boolean()
		}

		if element, err := p.Value().Document().LookupErr("exclusiveMinimum"); err == nil {
			exclusiveMinimum = element.Boolean()
		}

		if element, err := p.Value().Document().LookupErr("enum"); err == nil {
			rawEnums, err := element.Document().Elements()

			if err != nil {
				log.Println(err)
				return nil, err
			}
			enum = make([]interface{}, len(rawEnums))

			for i, e := range rawEnums {
				if e.Value().IsNumber() {
					enum[i] = e.Value().Decimal128()
				} else {
					enum[i] = e.Value().StringValue()
				}
			}
		}

		properties[i] = model.NewValidationProperty(
			p.Key(),
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
		)
	}

	// The validationLevel option determines which operations MongoDB applies the validation rules:
	// - If the validationLevel is strict (the default), MongoDB applies validation rules to all inserts and updates.
	// - f the validationLevel is moderate, MongoDB applies validation rules to inserts and to updates to existing documents that already fulfill the validation criteria. With the moderate level, updates to existing documents that do not fulfill the validation criteria are not checked for validity.
	var validationLevel value.ValidationLevel
	// The validationAction option determines how MongoDB handles documents that violate the validation rules:
	// - If the validationAction is error (the default), MongoDB rejects any insert or update that violates the validation criteria.
	// - If the validationAction is warn, MongoDB logs any violations but allows the insertion or update to proceed.
	var validationAction value.ValidationAction

	if element, err := options.Document().LookupErr("validationLevel"); err == nil {
		validationLevel = value.ValidationLevel(element.StringValue())
	}

	if element, err := options.Document().LookupErr("validationAction"); err == nil {
		validationAction = value.ValidationAction(element.StringValue())
	}

	return model.NewValidation(validationLevel, validationAction, properties), nil
}
