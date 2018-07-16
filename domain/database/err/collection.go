package err

import "errors"

var (
	EmptyCollName = errors.New("collection name should not be blank")
)
