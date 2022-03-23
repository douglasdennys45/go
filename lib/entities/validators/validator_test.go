package validators_test

import (
	"clean-architeture/lib/entities/validators"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakerLocal struct {
	Username string  `faker:"username"`
	Age      int64   `faker:"oneof: 1, 2"`
	Amount   float64 `faker:"amount"`
	Email    string  `faker:"email"`
}

func TestValidator_IsEmail(t *testing.T) {
	mock := fakerLocal{}
	_ = faker.FakeData(&mock)

	validator := validators.Validator{}
	isEmail := validator.IsEmail(mock.Email)
	assert.Equal(t, isEmail, true)
}

func TestValidator_IsString(t *testing.T) {
	mock := fakerLocal{}
	_ = faker.FakeData(&mock)

	validator := validators.Validator{}
	isEmail := validator.IsString(mock.Username)
	assert.Equal(t, isEmail, true)
}

func TestValidator_IsDecimal(t *testing.T) {
	mock := fakerLocal{}
	_ = faker.FakeData(&mock)

	validator := validators.Validator{}
	isEmail := validator.IsDecimal(mock.Amount)
	assert.Equal(t, isEmail, true)
}

func TestValidator_IsInt(t *testing.T) {
	mock := fakerLocal{}
	_ = faker.FakeData(&mock)

	validator := validators.Validator{}
	isEmail := validator.IsInt(mock.Age)
	assert.Equal(t, isEmail, true)
}
