package uuid_test

import (
	"github.com/douglasdennys45/go/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUUID(t *testing.T) {
	u := uuid.NewUUID()
	assert.NotNil(t, u)
}
