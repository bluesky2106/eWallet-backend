package main

import (
	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/integration_test/config"
)

func main() {
	conf := commonConfig.ParseConfig("config.json", "../config")
	testConf := config.ParseConfig(conf)
	testConf.Print()
}
