package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "admin")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	assert.NotPanics(t, func() {
		Connect()
	})
}
