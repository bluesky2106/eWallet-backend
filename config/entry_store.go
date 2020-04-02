package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// EntryStore configurations
type EntryStore struct {
	Host  string `json: "host"`
	Port  string `json: "port"`
	MySQL MySQL  `json: "mysql"`
}

func setDefaultEntryStore() {
	viper.SetDefault("entryStore.host", "localhost")
	viper.SetDefault("entryStore.port", "3")
	setDefaultMySQL("entryStore")
}

func (conf *Config) printEntryStoreConfig() {
	fmt.Println("--------- Entry Store configurations -----------")
	fmt.Println("Server host is\t", conf.EntryStore.Host)
	fmt.Println("Server port is\t", conf.EntryStore.Port)

	fmt.Println("MySQL DB name is\t", conf.EntryStore.MySQL.DBName)
	fmt.Println("MySQL User is\t\t", conf.EntryStore.MySQL.Username)
	fmt.Println("MySQL Pass is\t\t", conf.EntryStore.MySQL.Password)
	fmt.Println("MySQL Host is\t\t", conf.EntryStore.MySQL.Host)
	fmt.Println("MySQL Port is\t\t", conf.EntryStore.MySQL.Port)
}
