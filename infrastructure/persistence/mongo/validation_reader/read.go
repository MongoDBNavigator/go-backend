package validation_reader

import (
	"context"
	"log"

	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/mongodb/mongo-go-driver/bson"
)

// Read method to read $jsonSchema validation
// JSON Schema is the recommended means of performing schema validation.
// https://docs.mongodb.com/manual/reference/operator/query/jsonSchema/#jsonschema
// http://json-schema.org/
func (rcv *validationReader) Read(dbName value.DBName, collName value.CollName) (*model.Validation, error) {
	reader, err := rcv.db.Database(string(dbName)).RunCommand(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("listCollections", 1),
			bson.EC.SubDocument("filter", bson.NewDocument(
				bson.EC.String("name", string(collName)),
			)),
		),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cursor, err := reader.Lookup("cursor")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	firstBatch, err := cursor.Value().ReaderDocument().Lookup("firstBatch")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	data, err := firstBatch.Value().MutableArray().Lookup(0)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	options, err := data.ReaderDocument().Lookup("options")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	validator, err := options.Value().ReaderDocument().Lookup("validator")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var jsonSchema bson.Reader

	// MongoDB allowed validation:
	// - with $jsonSchema property
	// - without $jsonSchema property
	if element, err := validator.Value().ReaderDocument().Lookup("$jsonSchema"); err == nil {
		jsonSchema = element.Value().ReaderDocument()
	} else {
		jsonSchema = validator.Value().ReaderDocument()
	}

	var requiredFields []string

	// collect required fields
	if element, err := jsonSchema.Lookup("required"); err == nil {
		requiredFields = make([]string, element.Value().MutableArray().Len())
		iterator, err := element.Value().MutableArray().Iterator()

		if err != nil {
			log.Println(err)
			return nil, err
		}

		i := 0
		for iterator.Next() {
			requiredFields[i] = iterator.Value().StringValue()
			i++
		}
	}

	prop, err := jsonSchema.Lookup("properties")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	propIterator, err := prop.Value().ReaderDocument().Iterator()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	properties := make([]*model.ValidationProperty, 0)

	for propIterator.Next() {
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
			if v == propIterator.Element().Key() {
				required = true
				break
			}
		}

		elementReader := propIterator.Element().Value().ReaderDocument()

		if element, err := elementReader.Lookup("bsonType"); err == nil {
			bsonType = element.Value().StringValue()
		}

		if element, err := elementReader.Lookup("description"); err == nil {
			description = element.Value().StringValue()
		}

		if element, err := elementReader.Lookup("pattern"); err == nil {
			pattern = element.Value().StringValue()
		}

		if element, err := elementReader.Lookup("minimum"); err == nil {
			minimum = int(element.Value().Int32())
		}

		if element, err := elementReader.Lookup("maximum"); err == nil {
			maximum = int(element.Value().Int32())
		}

		if element, err := elementReader.Lookup("maxLength"); err == nil {
			maxLength = int(element.Value().Int32())
		}

		if element, err := elementReader.Lookup("minLength"); err == nil {
			minLength = int(element.Value().Int32())
		}

		if element, err := elementReader.Lookup("uniqueItems"); err == nil {
			uniqueItems = element.Value().Boolean()
		}

		if element, err := elementReader.Lookup("exclusiveMaximum"); err == nil {
			exclusiveMaximum = element.Value().Boolean()
		}

		if element, err := elementReader.Lookup("exclusiveMinimum"); err == nil {
			exclusiveMinimum = element.Value().Boolean()
		}

		if element, err := elementReader.Lookup("enum"); err == nil {
			enum = make([]interface{}, element.Value().MutableArray().Len())

			iterator, err := element.Value().MutableArray().Iterator()

			if err != nil {
				log.Println(err)
				return nil, err
			}

			i := 0
			for iterator.Next() {
				enum[i] = iterator.Value().Interface()
				i++
			}
		}

		properties = append(properties, model.NewValidationProperty(
			propIterator.Element().Key(),
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

	// The validationLevel option determines which operations MongoDB applies the validation rules:
	// - If the validationLevel is strict (the default), MongoDB applies validation rules to all inserts and updates.
	// - f the validationLevel is moderate, MongoDB applies validation rules to inserts and to updates to existing documents that already fulfill the validation criteria. With the moderate level, updates to existing documents that do not fulfill the validation criteria are not checked for validity.
	var validationLevel value.ValidationLevel
	// The validationAction option determines how MongoDB handles documents that violate the validation rules:
	// - If the validationAction is error (the default), MongoDB rejects any insert or update that violates the validation criteria.
	// - If the validationAction is warn, MongoDB logs any violations but allows the insertion or update to proceed.
	var validationAction value.ValidationAction

	if element, err := options.Value().ReaderDocument().Lookup("validationLevel"); err == nil {
		validationLevel = value.ValidationLevel(element.Value().StringValue())
	}

	if element, err := options.Value().ReaderDocument().Lookup("validationAction"); err == nil {
		validationAction = value.ValidationAction(element.Value().StringValue())
	}

	return model.NewValidation(validationLevel, validationAction, properties), nil
}
