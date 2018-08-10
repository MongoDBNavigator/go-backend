package err

import "errors"

var (
	EmptyCollName                = errors.New("collection name should not be blank")
	NotValidCollName             = errors.New(`collection name is not valid (pattern "/\. "$")`)
	SystemPrefixContainsCollName = errors.New(`collection name is not valid ('system.' prefix reserved for internal use)`)
)
