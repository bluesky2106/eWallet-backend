package config

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
)

// Config : integration test configurations
type Config struct {
	Env   string `json:"env"`
	URLBO string `json:"URLBO"`
	DbBO  string `json:"db_bo"`
}

// ParseConfig : parse configurations from common config
func ParseConfig(conf *config.Config) *Config {
	return &Config{
		Env:   string(conf.Env),
		URLBO: fmt.Sprintf("http://%s:%s/api", conf.BOController.Host, conf.BOController.Port),
		DbBO: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			conf.BOEntryStore.MySQL.Username,
			conf.BOEntryStore.MySQL.Password,
			conf.BOEntryStore.MySQL.Host,
			conf.BOEntryStore.MySQL.Port,
			conf.BOEntryStore.MySQL.DBName),
	}
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tEnv:\t\t\t%s\n", conf.Env)
	fmt.Printf("\t\tBO URL:\t\t\t%s\n", conf.URLBO)
	fmt.Printf("\t\tBO Database:\t\t%s\n", conf.DbBO)
}
