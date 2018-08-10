package response

import (
	"github.com/MongoDBNavigator/go-backend/domain/database/model"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/database_resource/representation"
)

func ValidationToView(model *model.Validation) *representation.Validation {
	properties := make([]*representation.ValidationProperty, len(model.Properties()))

	for i, prop := range model.Properties() {
		properties[i] = &representation.ValidationProperty{
			Name:             prop.Name(),
			Enum:             prop.Enum(),
			Pattern:          prop.Pattern(),
			Minimum:          prop.Minimum(),
			Maximum:          prop.Maximum(),
			MinLength:        prop.MinLength(),
			Type:             prop.BsonType(),
			Required:         prop.Required(),
			MaxLength:        prop.MaxLength(),
			Description:      prop.Description(),
			UniqueItems:      prop.UniqueItems(),
			ExclusiveMaximum: prop.ExclusiveMaximum(),
			ExclusiveMinimum: prop.ExclusiveMinimum(),
			MinItems:         prop.MinItems(),
			MaxItems:         prop.MaxItems(),
		}
	}

	return &representation.Validation{
		ValidationAction: string(model.ValidationAction()),
		ValidationLevel:  string(model.ValidationLevel()),
		Properties:       properties,
	}
}
