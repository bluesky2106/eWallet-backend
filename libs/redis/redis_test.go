package redis

import (
	"fmt"
	"testing"

	commonConfig "github.com/bluesky2106/eWallet-backend/config"

	"github.com/stretchr/testify/assert"
)

var conf *commonConfig.Config

func init() {
	conf = &commonConfig.Config{
		Redis: commonConfig.Redis{
			Host:     "localhost",
			Port:     "6379",
			DB:       0,
			Password: "3112",
		},
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
