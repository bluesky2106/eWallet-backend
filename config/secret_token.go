package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func setDefaultTokenSecretKey() {
	viper.SetDefault("tokenSecretKey", "")
}

func (conf *Config) printTokenSecretKey() {
	fmt.Println("--------- TokenSecretKey configurations -----------")
	fmt.Println("TokenSecretKey is\t", conf.TokenSecretKey)
}
