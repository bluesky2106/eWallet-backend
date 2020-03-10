package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Program configurations
type Program struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultProgram() {
	viper.SetDefault("program.host", "localhost")
	viper.SetDefault("program.port", "9")
}

func (conf *Config) printProgramConfig() {
	fmt.Println("----------- Program configurations -------------")
	fmt.Println("Server host is\t", conf.Program.Host)
	fmt.Println("Server port is\t", conf.Program.Port)
}
