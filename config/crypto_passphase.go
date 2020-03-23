package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func setDefaultCryptoPassphase() {
	viper.SetDefault("cryptoPassphase", "")
}

func (conf *Config) printCryptoPassphase() {
	fmt.Println("--------- CryptoPassphase configurations -----------")
	fmt.Println("CryptoPassphase is\t", conf.CryptoPassphase)
}
