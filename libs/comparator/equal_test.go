package comparator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {
	assert := assert.New(t)

	a := new(int)
	*a = 8
	b := new(int)
	*b = 8
	assert.True(IsVariableValueEqual(a, b))
}
