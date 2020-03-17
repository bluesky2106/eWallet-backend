package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Environment : debug or production mode
type Environment string

const (
	// Debug mode
	Debug Environment = "debug"
	// Production mode
	Production Environment = "production"
)

func setDefaultEnvironment() {
	viper.SetDefault("env", Debug)
}

func (conf *Config) printEnvironmentConfig() {
	fmt.Println("--------- Environment configurations -----------")
	fmt.Println("Environment is\t", conf.Env)
}
