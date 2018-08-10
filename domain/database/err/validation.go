package err

import "errors"

var (
	ValidationNotFound = errors.New("validation not found")

	EmptyValidationLevel   = errors.New("validation level should not be blank")
	InvalidValidationLevel = errors.New(`validation level value you selected is not a valid choice ("off", "moderate", "strict")`)

	EmptyValidationAction   = errors.New("validation action should not be blank")
	InvalidValidationAction = errors.New(`validation action value you selected is not a valid choice ("error", "warning")`)

	MinLengthGreatMaxLength = errors.New("maxLength should be greater than minLength")
	MinItemsGreatMaxItems   = errors.New("minItems should be greater than maxItems")
	MinimumGreatMaximum     = errors.New("minimum should be greater than maximum")
)
