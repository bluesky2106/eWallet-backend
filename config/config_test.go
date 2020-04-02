package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	assert := assert.New(t)

	conf := ParseConfig("config.json", "./")
	assert.NotNil(conf)

	conf.Print()
}
