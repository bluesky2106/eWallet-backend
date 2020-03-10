package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// EntryScan configurations
type EntryScan struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultEntryScan() {
	viper.SetDefault("entryScan.host", "localhost")
	viper.SetDefault("entryScan.port", "5")
}

func (conf *Config) printEntryScanConfig() {
	fmt.Println("---------- Entry Scan configurations -----------")
	fmt.Println("Server host is\t", conf.EntryScan.Host)
	fmt.Println("Server port is\t", conf.EntryScan.Port)
}
