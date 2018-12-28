package validator_writer

import (
	"regexp"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"

	localErr "github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
)

// build structure for $jsonSchema
func buildJsonSchema(validation *model.Validation) (*bson.Document, error) {
	jsonSchema := bson.NewDocument()
	jsonSchema.Set(bson.EC.String("bsonType", "object"))

	requiredFields := make([]*bson.Value, 0)
	propertiesDocument := bson.NewDocument()

	for _, prop := range validation.Properties() {
		if prop.Required() {
			requiredFields = append(requiredFields, bson.VC.String(prop.Name()))
		}

		propertyDocument := bson.NewDocument()

		if prop.Description() != "" {
			propertyDocument.Set(bson.EC.String("description", prop.Description()))
		}

		if len(prop.Enum()) > 0 {
			enumItems := make([]*bson.Value, len(prop.Enum()))
			for i, item := range prop.Enum() {
				enumItems[i] = bson.EC.Interface(strconv.Itoa(i), item).Value()
			}

			propertyDocument.Set(bson.EC.Array("enum", bson.NewArray(enumItems...)))
		} else {
			propertyDocument.Set(bson.EC.String("bsonType", prop.BsonType()))

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

		propertiesDocument.Set(bson.EC.SubDocument(prop.Name(), propertyDocument))
	}

	jsonSchema.Set(bson.EC.Array("required", bson.NewArray(requiredFields...)))
	jsonSchema.Set(bson.EC.SubDocument("properties", propertiesDocument))

	document := bson.NewDocument()
	document.Set(bson.EC.SubDocument("$jsonSchema", jsonSchema))

	return document, nil
}

// build string $jsonSchema validator
func buildStringValidator(prop *model.ValidationProperty, propValidator *bson.Document) error {
	if prop.Pattern() != "" {
		if _, err := regexp.Compile(prop.Pattern()); err != nil {
			return err
		}
		propValidator.Set(bson.EC.String("pattern", prop.Pattern()))
	}

	if prop.MinLength() != 0 && prop.MaxLength() != 0 && prop.MinLength() > prop.MaxLength() {
		return localErr.MinLengthGreatMaxLength
	}

	if prop.MinLength() != 0 {
		propValidator.Set(bson.EC.Int32("minLength", int32(prop.MinLength())))
	}

	if prop.MaxLength() != 0 {
		propValidator.Set(bson.EC.Int32("maxLength", int32(prop.MaxLength())))
	}

	return nil
}

// build numbers $jsonSchema validator
func buildNumberValidator(prop *model.ValidationProperty, propValidator *bson.Document) error {
	if prop.Minimum() != 0 && prop.Maximum() != 0 && prop.Minimum() > prop.Maximum() {
		return localErr.MinimumGreatMaximum
	}

	if prop.Minimum() != 0 {
		propValidator.Set(bson.EC.Int32("minimum", int32(prop.Minimum())))
		propValidator.Set(bson.EC.Boolean("exclusiveMinimum", prop.ExclusiveMinimum()))
	}

	if prop.Maximum() != 0 {
		propValidator.Set(bson.EC.Int32("maximum", int32(prop.Maximum())))
		propValidator.Set(bson.EC.Boolean("exclusiveMaximum", prop.ExclusiveMaximum()))
	}

	return nil
}

// build arrays $jsonSchema validator
func buildArrayValidator(prop *model.ValidationProperty, propValidator *bson.Document) error {
	if prop.MinItems() != 0 && prop.MaxItems() != 0 && prop.MinItems() > prop.MaxItems() {
		return localErr.MinItemsGreatMaxItems
	}

	if prop.MinItems() != 0 {
		propValidator.Set(bson.EC.Int32("minItems", int32(prop.MinItems())))
	}

	if prop.MaxItems() != 0 {
		propValidator.Set(bson.EC.Int32("maxItems", int32(prop.MaxItems())))
	}

	propValidator.Set(bson.EC.Boolean("uniqueItems", prop.UniqueItems()))

	return nil
}
