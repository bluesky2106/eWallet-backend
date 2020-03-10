package config

import "github.com/spf13/viper"

// EntryCache configurations
type EntryCache struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultEntryCache() {
	viper.SetDefault("entryCache.host", "localhost")
	viper.SetDefault("entryCache.port", "2")
}
