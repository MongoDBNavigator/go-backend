package model

import "github.com/MongoDBNavigator/go-backend/domain/database/value"

// Validation model
type Validation struct {
	validationAction value.ValidationAction
	validationLevel  value.ValidationLevel
	properties       []*ValidationProperty
}

// Getter for properties
func (v *Validation) Properties() []*ValidationProperty {
	return v.properties
}

// Getter for validationLevel
func (v *Validation) ValidationLevel() value.ValidationLevel {
	return v.validationLevel
}

// Getter for validationAction
func (v *Validation) ValidationAction() value.ValidationAction {
	return v.validationAction
}

// Validation Property (field name and constraints)
type ValidationProperty struct {
	name             string
	required         bool
	bsonType         string
	enum             []interface{}
	description      string
	minimum          int
	maximum          int
	exclusiveMaximum bool
	exclusiveMinimum bool
	uniqueItems      bool
	pattern          string
	maxLength        int
	minLength        int
	minItems         int
	maxItems         int
}

// Getter for maxItems
func (v *ValidationProperty) MaxItems() int {
	return v.maxItems
}

// Getter for minItems
func (v *ValidationProperty) MinItems() int {
	return v.minItems
}

// Getter for exclusiveMinimum
func (v *ValidationProperty) ExclusiveMinimum() bool {
	return v.exclusiveMinimum
}

// Getter for exclusiveMaximum
func (v *ValidationProperty) ExclusiveMaximum() bool {
	return v.exclusiveMaximum
}

// Getter for uniqueItems
func (v *ValidationProperty) UniqueItems() bool {
	return v.uniqueItems
}

// Getter for minLength
func (v *ValidationProperty) MinLength() int {
	return v.minLength
}

// Getter for maxLength
func (v *ValidationProperty) MaxLength() int {
	return v.maxLength
}

// Getter for pattern
func (v *ValidationProperty) Pattern() string {
	return v.pattern
}

// Getter for maximum
func (v *ValidationProperty) Maximum() int {
	return v.maximum
}

// Getter for minimum
func (v *ValidationProperty) Minimum() int {
	return v.minimum
}

// Getter for description
func (v *ValidationProperty) Description() string {
	return v.description
}

// Getter for enum
func (v *ValidationProperty) Enum() []interface{} {
	return v.enum
}

// Getter for bsonType
func (v *ValidationProperty) BsonType() string {
	return v.bsonType
}

// Getter for required
func (v *ValidationProperty) Required() bool {
	return v.required
}

// Getter for name
func (v *ValidationProperty) Name() string {
	return v.name
}

func NewValidation(
	validationLevel value.ValidationLevel,
	validationAction value.ValidationAction,
	properties []*ValidationProperty,
) *Validation {
	return &Validation{
		validationLevel:  validationLevel,
		validationAction: validationAction,
		properties:       properties,
	}
}

func NewValidationProperty(
	name string,
	required bool,
	bsonType string,
	enum []interface{},
	description string,
	minimum int,
	maximum int,
	pattern string,
	maxLength int,
	minLength int,
	minItems int,
	maxItems int,
	exclusiveMaximum bool,
	exclusiveMinimum bool,
	uniqueItems bool,
) *ValidationProperty {
	return &ValidationProperty{
		name:             name,
		required:         required,
		bsonType:         bsonType,
		enum:             enum,
		description:      description,
		minimum:          minimum,
		maximum:          maximum,
		pattern:          pattern,
		maxLength:        maxLength,
		minLength:        minLength,
		minItems:         minItems,
		maxItems:         maxItems,
		exclusiveMaximum: exclusiveMaximum,
		exclusiveMinimum: exclusiveMinimum,
		uniqueItems:      uniqueItems,
	}
}
