package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// RabbitMQ configurations
type RabbitMQ struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultRabbitMQ() {
	viper.SetDefault("rabbitmq.username", "rabbitmq")
	viper.SetDefault("rabbitmq.password", "")
	viper.SetDefault("rabbitmq.host", "localhost")
	viper.SetDefault("rabbitmq.port", "5672")
}

func (conf *Config) printRabbitMQConfig() {
	fmt.Println("------------ RabbitMQ configurations -----------")
	fmt.Println("RabbitMQ User is\t", conf.RabbitMQ.Username)
	fmt.Println("RabbitMQ Pass is\t", conf.RabbitMQ.Password)
	fmt.Println("RabbitMQ Host is\t", conf.RabbitMQ.Host)
	fmt.Println("RabbitMQ Port is\t", conf.RabbitMQ.Port)
}
