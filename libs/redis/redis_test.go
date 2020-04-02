package redis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var conf *Config

func init() {
	conf = &Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
}

func TestInit(t *testing.T) {
	assert := assert.New(t)

	// init redis
	client, err := Init(conf)

	if err != nil {
		fmt.Errorf("failed to connect redis server: %v", err)
	}
	assert.Nil(err)
	assert.NotNil(client)
}
