package config_test

import (
	"github.com/douglasdennys45/go/internal/infrastructure/config"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestRouterPostUsers(t *testing.T) {
	app := config.NewServer()
	response, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/users", nil))
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestRouterGetUsers(t *testing.T) {
	app := config.NewServer()
	response, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/users/1", nil))
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestRouterNotFound(t *testing.T) {
	app := config.NewServer()
	response, err := app.Test(httptest.NewRequest(fiber.MethodPatch, "/users/1", nil))
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.StatusCode, 404)
}
