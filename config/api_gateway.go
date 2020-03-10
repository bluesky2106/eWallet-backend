package config

import "github.com/spf13/viper"

// APIGateway configurations
type APIGateway struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultAPIGateway() {
	viper.SetDefault("apiGateway.host", "localhost")
	viper.SetDefault("apiGateway.port", "1")
}
