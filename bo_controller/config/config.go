package config

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/libs/rabbitmq"
)

// Config : BO Controller configuration
type Config struct {
	Host           string `json: "host"`
	Port           string `json: "port"`
	Env            string `json: "env"`
	TokenSecretKey string `json:"tokenSecretKey"`

	// rabbit
	RabbitMQ *rabbitmq.Config

	// services
	EntryStoreEndpoint   string `json:"entryStoreEndpoint"`
	BOEntryStoreEndpoint string `json:"boEntryStoreEndpoint"`
}

// ParseConfig : get configurations related to bo controller from common configurations
func ParseConfig(conf *config.Config) *Config {
	return &Config{
		Host:           conf.BOController.Host,
		Port:           conf.BOController.Port,
		Env:            string(conf.Env),
		TokenSecretKey: conf.TokenSecretKey,
		RabbitMQ: &rabbitmq.Config{
			URL:      fmt.Sprintf("%s:%s", conf.RabbitMQ.Host, conf.RabbitMQ.Port),
			User:     conf.RabbitMQ.Username,
			Password: conf.RabbitMQ.Password,
		},
		EntryStoreEndpoint:   fmt.Sprintf("%s:%s", conf.EntryStore.Host, conf.EntryStore.Port),
		BOEntryStoreEndpoint: fmt.Sprintf("%s:%s", conf.BOEntryStore.Host, conf.BOEntryStore.Port),
	}
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tEnv:\t\t\t%s\n", conf.Env)
	fmt.Printf("\t\tHost:\t\t\t%s\n", conf.Host)
	fmt.Printf("\t\tPort:\t\t\t%s\n", conf.Port)
	fmt.Printf("\t\tTokenSecretKey:\t\t%s\n", conf.TokenSecretKey)
	fmt.Printf("\t\tRabbitMQURL:\t\t%s\n", conf.RabbitMQ.URL)
	fmt.Printf("\t\tRabbitMQUser:\t\t%s\n", conf.RabbitMQ.User)
	fmt.Printf("\t\tRabbitMQPwd:\t\t%s\n", conf.RabbitMQ.Password)
	fmt.Printf("\t\tEntryStoreEndpoint:\t%s\n", conf.EntryStoreEndpoint)
	fmt.Printf("\t\tBOEntryStoreEndpoint:\t%s\n", conf.BOEntryStoreEndpoint)
}
