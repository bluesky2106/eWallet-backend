package config

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
)

// Config : BO Controller configuration
type Config struct {
	Host           string `json: "host"`
	Port           string `json: "port"`
	Env            string `json: "env"`
	TokenSecretKey string `json:"tokenSecretKey"`

	// services
	BOEntryStoreEndpoint string `json:"boEntryStoreEndpoint"`
}

// ParseConfig : get configurations related to bo controller from common configurations
func ParseConfig(conf *config.Config) *Config {
	return &Config{
		Host:                 conf.BOController.Host,
		Port:                 conf.BOController.Port,
		Env:                  string(conf.Env),
		TokenSecretKey:       conf.TokenSecretKey,
		BOEntryStoreEndpoint: fmt.Sprintf("%s:%s", conf.BOEntryStore.Host, conf.BOEntryStore.Port),
	}
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tEnv:\t\t\t%s\n", conf.Env)
	fmt.Printf("\t\tHost:\t\t\t%s\n", conf.Host)
	fmt.Printf("\t\tPort:\t\t\t%s\n", conf.Port)
	fmt.Printf("\t\tTokenSecretKey:\t\t%s\n", conf.TokenSecretKey)
	fmt.Printf("\t\tBOEntryStoreEndpoint:\t%s\n", conf.BOEntryStoreEndpoint)
}
