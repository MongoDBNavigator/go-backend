package err

import "errors"

var (
	EmptyDBName            = errors.New("database name should not be blank")
	NotValidDBName         = errors.New(`database name is not valid (pattern "/\. "$")`)
	InvalidDBNameMaxLength = errors.New(`database name is too long (it should have 64 characters or less)`)
)
