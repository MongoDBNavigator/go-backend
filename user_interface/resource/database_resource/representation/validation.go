package representation

type Validation struct {
	ValidationAction string                `json:"validationAction"`
	ValidationLevel  string                `json:"validationLevel"`
	Properties       []*ValidationProperty `json:"properties"`
}

// Structure to representante for Validation Property
type ValidationProperty struct {
	Name             string        `json:"name"`
	Required         bool          `json:"required"`
	Type             string        `json:"type,omitempty"`
	Enum             []interface{} `json:"enum,omitempty"`
	Description      string        `json:"description,omitempty"`
	Minimum          int           `json:"minimum,omitempty"`
	Maximum          int           `json:"maximum,omitempty"`
	Pattern          string        `json:"pattern,omitempty"`
	MaxLength        int           `json:"maxLength,omitempty"`
	MinLength        int           `json:"minLength,omitempty"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum,omitempty"`
	UniqueItems      bool          `json:"uniqueItems,omitempty"`
	MinItems         int           `json:"minItems,omitempty"`
	MaxItems         int           `json:"maxItems,omitempty"`
}
