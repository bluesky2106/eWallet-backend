package main

import (
	"github.com/bluesky2106/eWallet-backend/config"
)

func main() {
	conf := config.ParseConfig("config.json", "config")
	conf.Print()
}
