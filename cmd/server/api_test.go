package main_test

import (
	"github.com/douglasdennys45/go/internal/infrastructure/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStart(t *testing.T) {
	app := config.NewServer()
	assert.NotNil(t, app)
}
