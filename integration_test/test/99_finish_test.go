package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClearAll(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(clearAdminUsers())
}
