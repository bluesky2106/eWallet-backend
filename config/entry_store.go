package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// EntryStore configurations
type EntryStore struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultEntryStore() {
	viper.SetDefault("entryStore.host", "localhost")
	viper.SetDefault("entryStore.port", "3")
}

func (conf *Config) printEntryStoreConfig() {
	fmt.Println("--------- Entry Store configurations -----------")
	fmt.Println("Server host is\t", conf.EntryStore.Host)
	fmt.Println("Server port is\t", conf.EntryStore.Port)
}
