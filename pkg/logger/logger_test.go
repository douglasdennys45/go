package logger_test

import (
	"github.com/douglasdennys45/go/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	log := logger.NewLogger()
	assert.NotNil(t, log)
}
