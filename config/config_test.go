package config

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	conf := ParseConfig("config.json", "./")
	conf.Print()
}
