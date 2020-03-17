package mysql

import (
	"log"
	"testing"

	"github.com/bluesky2106/eWallet-backend/config"
	"github.com/stretchr/testify/assert"
)

func TestInitMySQL(t *testing.T) {
	assert := assert.New(t)

	// load config from environments
	conf := &config.Config{
		MySQL: config.MySQL{
			Host:     "localhost",
			Port:     "3307",
			Username: "tokoin",
			Password: "tokoin",
			DBName:   "tokoin",
		},
		Env: config.Debug,
	}

	// init mysql
	dao, err := New(conf)

	if err != nil {
		log.Fatalf("failed to connect mysql: %v", err)
	}
	assert.Nil(err)
	assert.NotNil(dao)
}
