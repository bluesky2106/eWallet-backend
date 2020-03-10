package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Email configurations
type Email struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultEmail() {
	viper.SetDefault("email.host", "localhost")
	viper.SetDefault("email.port", "8")
}

func (conf *Config) printEmailConfig() {
	fmt.Println("------------ Email configurations --------------")
	fmt.Println("Server host is\t", conf.Email.Host)
	fmt.Println("Server port is\t", conf.Email.Port)
}
