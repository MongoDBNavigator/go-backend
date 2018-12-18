package helper

import (
	"log"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestGetVar(t *testing.T) {
	defaultValue := "--TEST--"
	variableName := "OS_ENV_TEST_VARIABLE"

	assert.Equal(t, defaultValue, GetVar(variableName, defaultValue))

	originValue := "SUPPER+TEST"
	if err := os.Setenv(variableName, originValue); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, originValue, GetVar(variableName, defaultValue))
}
