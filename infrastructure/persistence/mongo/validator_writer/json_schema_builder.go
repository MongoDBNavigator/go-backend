package validator_writer

import (
	"log"
	"regexp"

	"github.com/mongodb/mongo-go-driver/bson"

	localErr "github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
)

// build bson structure for $jsonSchema
func buildJsonSchema(validation *model.Validation) (*bson.D, error) {
	jsonSchema := bson.D{{"bsonType", "object"}}
	requiredFields := bson.A{}
	propertiesDocument := bson.D{}

	for _, prop := range validation.Properties() {
		if prop.Required() {
			requiredFields = append(requiredFields, prop.Name())
		}

		propertyDocument := bson.D{}
		if prop.Description() != "" {
			propertyDocument = append(propertyDocument, bson.E{Key: "description", Value: prop.Description()})
		}

		if len(prop.Enum()) > 0 {
			enumItems := bson.A{}
			for _, item := range prop.Enum() {
				enumItems = append(enumItems, item)
			}

			propertyDocument = append(propertyDocument, bson.E{Key: "enum", Value: enumItems})
		} else {
			propertyDocument = append(propertyDocument, bson.E{Key: "bsonType", Value: prop.BsonType()})

			switch prop.BsonType() {
			case "string":
				if err := buildStringValidator(prop, propertyDocument); err != nil {
					return nil, err
				}
			case "array":
				if err := buildArrayValidator(prop, propertyDocument); err != nil {
					return nil, err
				}
			case "double", "int", "long", "decimal":
				if err := buildNumberValidator(prop, propertyDocument); err != nil {
					return nil, err
				}
			}
		}

		propertiesDocument = append(propertiesDocument, bson.E{Key: prop.Name(), Value: propertyDocument})
	}

	if len(requiredFields) > 0 {
		jsonSchema = append(jsonSchema, bson.E{Key: "required", Value: requiredFields})
	}

	jsonSchema = append(jsonSchema, bson.E{Key: "properties", Value: propertiesDocument})

	log.Println(jsonSchema)

	return &bson.D{{"$jsonSchema", jsonSchema}}, nil
}

// build string $jsonSchema validator
func buildStringValidator(prop *model.ValidationProperty, propValidator bson.D) error {
	if prop.Pattern() != "" {
		if _, err := regexp.Compile(prop.Pattern()); err != nil {
			return err
		}

		propValidator = append(propValidator, bson.E{Key: "pattern", Value: prop.Pattern()})
	}

	if prop.MinLength() != 0 && prop.MaxLength() != 0 && prop.MinLength() > prop.MaxLength() {
		return localErr.MinLengthGreatMaxLength
	}

	if prop.MinLength() != 0 {
		propValidator = append(propValidator, bson.E{Key: "minLength", Value: int32(prop.MinLength())})
	}

	if prop.MaxLength() != 0 {
		propValidator = append(propValidator, bson.E{Key: "maxLength", Value: int32(prop.MaxLength())})
	}

	return nil
}

// build numbers $jsonSchema validator
func buildNumberValidator(prop *model.ValidationProperty, propValidator bson.D) error {
	if prop.Minimum() != 0 && prop.Maximum() != 0 && prop.Minimum() > prop.Maximum() {
		return localErr.MinimumGreatMaximum
	}

	if prop.Minimum() != 0 {
		propValidator = append(propValidator, bson.E{Key: "minimum", Value: int32(prop.Minimum())})
		propValidator = append(propValidator, bson.E{Key: "exclusiveMinimum", Value: prop.ExclusiveMinimum()})
	}

	if prop.Maximum() != 0 {
		propValidator = append(propValidator, bson.E{Key: "maximum", Value: int32(prop.Maximum())})
		propValidator = append(propValidator, bson.E{Key: "exclusiveMaximum", Value: prop.ExclusiveMaximum()})
	}

	return nil
}

// build arrays $jsonSchema validator
func buildArrayValidator(prop *model.ValidationProperty, propValidator bson.D) error {
	if prop.MinItems() != 0 && prop.MaxItems() != 0 && prop.MinItems() > prop.MaxItems() {
		return localErr.MinItemsGreatMaxItems
	}

	if prop.MinItems() != 0 {
		propValidator = append(propValidator, bson.E{Key: "minItems", Value: int32(prop.MinItems())})
	}

	if prop.MaxItems() != 0 {
		propValidator = append(propValidator, bson.E{Key: "maxItems", Value: int32(prop.MaxItems())})
	}

	propValidator = append(propValidator, bson.E{Key: "uniqueItems", Value: prop.UniqueItems()})

	return nil
}
