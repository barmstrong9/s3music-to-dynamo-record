package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	testVar := "Testing"

	assert.Equal(t, GetEnv("TEST_ENV_VAR", "Fallback"), "Fallback")

	os.Setenv("TEST_ENV_VAR", testVar)
	assert.Equal(t, GetEnv("TEST_ENV_VAR", "Fallback"), testVar)

}
