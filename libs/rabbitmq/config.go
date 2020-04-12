package rabbitmq

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
)

// Config : rabbitmq
type Config struct {
	URL      string
	User     string
	Password string
}

// ParseConfig : get configurations related to rabbitmq from common configurations
func ParseConfig(conf *config.Config) *Config {
	return &Config{
		URL:      conf.RabbitMQ.Host + ":" + conf.RabbitMQ.Port,
		User:     conf.RabbitMQ.Username,
		Password: conf.RabbitMQ.Password,
	}
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tURL:\t\t\t%s\n", conf.URL)
	fmt.Printf("\t\tUser:\t\t\t%s\n", conf.User)
	fmt.Printf("\t\tPassword:\t\t\t\t%s\n", conf.Password)
}
