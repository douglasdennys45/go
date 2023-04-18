package entities_test

import (
	"github.com/douglasdennys45/go/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := entities.NewUser("Douglas", "douglasdennys45@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Douglas", user.Name)
	assert.Equal(t, "douglasdennys45@gmail.com", user.Email)
	assert.Equal(t, "123456", user.Password)
}
