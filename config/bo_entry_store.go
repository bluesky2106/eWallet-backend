package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// BOEntryStore configurations
type BOEntryStore struct {
	Host  string `json: "host"`
	Port  string `json: "port"`
	MySQL MySQL  `json: "mysql"`
}

func setDefaultBOEntryStore() {
	viper.SetDefault("boEntryStore.host", "localhost")
	viper.SetDefault("boEntryStore.port", "10")
	setDefaultMySQL("boEntryStore")
}

func (conf *Config) printBOEntryStoreConfig() {
	fmt.Println("-------- BO Entry Store configurations ---------")
	fmt.Println("Server host is\t", conf.BOEntryStore.Host)
	fmt.Println("Server port is\t", conf.BOEntryStore.Port)

	fmt.Println("MySQL DB name is\t", conf.BOEntryStore.MySQL.DBName)
	fmt.Println("MySQL User is\t\t", conf.BOEntryStore.MySQL.Username)
	fmt.Println("MySQL Pass is\t\t", conf.BOEntryStore.MySQL.Password)
	fmt.Println("MySQL Host is\t\t", conf.BOEntryStore.MySQL.Host)
	fmt.Println("MySQL Port is\t\t", conf.BOEntryStore.MySQL.Port)
}
