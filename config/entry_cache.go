package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// EntryCache configurations
type EntryCache struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultEntryCache() {
	viper.SetDefault("entryCache.host", "localhost")
	viper.SetDefault("entryCache.port", "2")
}

func (conf *Config) printEntryCacheConfig() {
	fmt.Println("--------- Entry Cache configurations -----------")
	fmt.Println("Server host is\t", conf.EntryCache.Host)
	fmt.Println("Server port is\t", conf.EntryCache.Port)
}
