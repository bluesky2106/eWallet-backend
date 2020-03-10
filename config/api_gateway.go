package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// APIGateway configurations
type APIGateway struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultAPIGateway() {
	viper.SetDefault("apiGateway.host", "localhost")
	viper.SetDefault("apiGateway.port", "1")
}

func (conf *Config) printAPIGatewayConfig() {
	fmt.Println("--------- API Gateway configurations -----------")
	fmt.Println("Server host is\t", conf.APIGateway.Host)
	fmt.Println("Server port is\t", conf.APIGateway.Port)
}
