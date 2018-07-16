package err

import "errors"

var (
	EmptyIndexName = errors.New("index name should not be blank")
)
