package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Event configurations
type Event struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultEvent() {
	viper.SetDefault("event.host", "localhost")
	viper.SetDefault("event.port", "6")
}

func (conf *Config) printEventConfig() {
	fmt.Println("------------ Event configurations --------------")
	fmt.Println("Server host is\t", conf.Event.Host)
	fmt.Println("Server port is\t", conf.Event.Port)
}
