package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Asset configurations
type Asset struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultAsset() {
	viper.SetDefault("asset.host", "localhost")
	viper.SetDefault("asset.port", "4")
}

func (conf *Config) printAssetConfig() {
	fmt.Println("------------ Asset configurations --------------")
	fmt.Println("Server host is\t", conf.Asset.Host)
	fmt.Println("Server port is\t", conf.Asset.Port)
}
