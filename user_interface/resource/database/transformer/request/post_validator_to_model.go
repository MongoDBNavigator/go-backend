package request

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/domain/database/value"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database/representation"
)

func PostValidatorConvertToModel(postRequest *representation.Validation) *model.Validation {
	properties := make([]*model.ValidationProperty, len(postRequest.Properties))

	for i, prop := range postRequest.Properties {
		properties[i] = model.NewValidationProperty(
			prop.Name,
			prop.Required,
			prop.Type,
			prop.Enum,
			prop.Description,
			prop.Minimum,
			prop.Maximum,
			prop.Pattern,
			prop.MaxLength,
			prop.MinLength,
			prop.MinItems,
			prop.MaxItems,
			prop.ExclusiveMaximum,
			prop.ExclusiveMinimum,
			prop.UniqueItems,
		)
	}

	return model.NewValidation(
		value.ValidationLevel(postRequest.ValidationLevel),
		value.ValidationAction(postRequest.ValidationAction),
		properties,
	)
}
