package helper

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestGetVar(t *testing.T) {
	defaultValue := "--TEST--"
	variableName := "OS_ENV_TEST_VARIABLE"

	assert.Equal(t, defaultValue, GetVar(variableName, defaultValue))

	originValue := "SUPPER+TEST"
	os.Setenv(variableName, originValue)

	assert.Equal(t, originValue, GetVar(variableName, defaultValue))
}
