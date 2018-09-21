package system_info_reader

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MongoDBNavigator/go-backend/domain/system/repository"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*repository.SystemInfoReader)(nil), new(systemInfoReader))
}
