package env_test

import (
	"github.com/douglasdennys45/go/pkg/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEnv(t *testing.T) {
	envs, err := env.NewEnv(".")
	assert.Nil(t, err)
	assert.NotNil(t, envs)
}
