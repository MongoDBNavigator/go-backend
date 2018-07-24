package representation

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Structure to representante POST body for Validation
type PostValidation struct {
	Properties       []*PostValidationProperty `json:"properties"`
	ValidationLevel  value.ValidationLevel     `json:"validationLevel"`
	ValidationAction value.ValidationAction    `json:"validationAction"`
}

// Structure to representante POST body for Validation Property
type PostValidationProperty struct {
	Name        string   `json:"name"`
	Required    bool     `json:"required"`
	Type        string   `json:"type,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Description string   `json:"description,omitempty"`
	Minimum     int      `json:"minimum,omitempty"`
	Maximum     int      `json:"maximum,omitempty"`
	Pattern     string   `json:"pattern,omitempty"`
	MaxLength   int      `json:"maxLength,omitempty"`
	MinLength   int      `json:"minLength,omitempty"`
}
