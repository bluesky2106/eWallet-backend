package config

import "github.com/spf13/viper"

// EntryStore configurations
type EntryStore struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultEntryStore() {
	viper.SetDefault("entryStore.host", "localhost")
	viper.SetDefault("entryStore.port", "3")
}
