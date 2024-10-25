package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	assert.NotPanics(t, func() {
		Connect()
	})
}
