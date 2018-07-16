package err

import "errors"

var (
	EmptyDocId = errors.New("document ID should not be blank")
)
