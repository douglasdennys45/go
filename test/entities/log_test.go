package entities_test

import (
	"clean-architeture/lib/entities"
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakerLocal struct {
	Type    string  `faker:"name"`
	Service string  `faker:"name"`
	Version float64 `faker:"amount"`
}

func convertToByte(data interface{}) []byte {
	convert, _ := json.Marshal(data)
	return convert
}

func TestLog_Create(t *testing.T) {
	mock := fakerLocal{}
	_ = faker.FakeData(&mock)

	isCreate, err := entities.Create(convertToByte(mock))
	assert.Nil(t, err)
	assert.NotNil(t, isCreate)
}

func TestLog_CreateRequired(t *testing.T) {
	mock := fakerLocal{Type: "", Service: "", Version: 0}

	_, err := entities.Create(convertToByte(mock))
	assert.NotNil(t, err)
}
