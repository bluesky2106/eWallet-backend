package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// BOController configurations
type BOController struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultBOController() {
	viper.SetDefault("boController.host", "localhost")
	viper.SetDefault("boController.port", "10")
}

func (conf *Config) printBOControllerConfig() {
	fmt.Println("-------- BO Controller configurations ----------")
	fmt.Println("Server host is\t", conf.BOController.Host)
	fmt.Println("Server port is\t", conf.BOController.Port)
}
