package errors_test

import (
	"clean-architeture/lib/entities/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_GenerateError(t *testing.T) {
	generate := errors.GenerateError("NOT_FOUND", 404, "NOT_FOUND", "NOT_FOUND")
	assert.NotNil(t, generate)
	assert.Equal(t, generate.Code, "NOT_FOUND")
	assert.Equal(t, generate.Title, "NOT_FOUND")
	assert.Equal(t, generate.Detail, "NOT_FOUND")
	assert.Equal(t, generate.Status, 404)
}
