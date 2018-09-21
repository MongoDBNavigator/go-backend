package validator_writer

import (
	"regexp"

	localErr "github.com/MongoDBNavigator/go-backend/domain/database/err"
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"gopkg.in/mgo.v2/bson"
)

// build structure for $jsonSchema
func buildJsonSchema(validation *model.Validation) (bson.M, error) {
	properties := bson.M{}
	requiredProps := make([]string, 0)
	jsonSchema := bson.M{
		"bsonType": "object",
	}

	for _, prop := range validation.Properties() {
		if prop.Required() {
			requiredProps = append(requiredProps, prop.Name())
		}

		propValidator := bson.M{}

		if prop.Description() != "" {
			propValidator["description"] = prop.Description()
		}

		if len(prop.Enum()) > 0 {
			propValidator["enum"] = prop.Enum()
		} else {
			propValidator["bsonType"] = prop.BsonType()

			switch prop.BsonType() {
			case "string":
				if err := buildStringValidator(prop, propValidator); err != nil {
					return nil, err
				}
			case "array":
				if err := buildArrayValidator(prop, propValidator); err != nil {
					return nil, err
				}
			case "double", "int", "long", "decimal":
				if err := buildNumberValidator(prop, propValidator); err != nil {
					return nil, err
				}
			}
		}

		properties[prop.Name()] = propValidator
	}

	if len(requiredProps) > 0 {
		jsonSchema["required"] = requiredProps
	}

	if len(properties) > 0 {
		jsonSchema["properties"] = properties
	}

	return jsonSchema, nil
}

// build string $jsonSchema validator
func buildStringValidator(prop *model.ValidationProperty, propValidator bson.M) error {
	if prop.Pattern() != "" {
		if _, err := regexp.Compile(prop.Pattern()); err != nil {
			return err
		}
		propValidator["pattern"] = prop.Pattern()
	}

	if prop.MinLength() != 0 && prop.MaxLength() != 0 && prop.MinLength() > prop.MaxLength() {
		return localErr.MinLengthGreatMaxLength
	}

	if prop.MinLength() != 0 {
		propValidator["minLength"] = prop.MinLength()
	}

	if prop.MaxLength() != 0 {
		propValidator["maxLength"] = prop.MaxLength()
	}

	return nil
}

// build numbers $jsonSchema validator
func buildNumberValidator(prop *model.ValidationProperty, propValidator bson.M) error {
	if prop.Minimum() != 0 && prop.Maximum() != 0 && prop.Minimum() > prop.Maximum() {
		return localErr.MinimumGreatMaximum
	}

	if prop.Minimum() != 0 {
		propValidator["minimum"] = prop.Minimum()
		propValidator["exclusiveMinimum"] = prop.ExclusiveMinimum()
	}

	if prop.Maximum() != 0 {
		propValidator["maximum"] = prop.Maximum()
		propValidator["exclusiveMaximum"] = prop.ExclusiveMaximum()
	}

	return nil
}

// build arrays $jsonSchema validator
func buildArrayValidator(prop *model.ValidationProperty, propValidator bson.M) error {
	if prop.MinItems() != 0 && prop.MaxItems() != 0 && prop.MinItems() > prop.MaxItems() {
		return localErr.MinItemsGreatMaxItems
	}

	if prop.MinItems() != 0 {
		propValidator["minItems"] = prop.MinItems()
	}

	if prop.MaxItems() != 0 {
		propValidator["maxItems"] = prop.MaxItems()
	}

	propValidator["uniqueItems"] = prop.UniqueItems()

	return nil
}
