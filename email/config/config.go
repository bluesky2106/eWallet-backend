package config

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/libs/rabbitmq"
	"github.com/bluesky2106/eWallet-backend/libs/sendgrid"
)

// Config : email service configurations
type Config struct {
	// env
	Host               string `json: "host"`
	Port               string `json: "port"`
	Env                string `json: "env"`
	EntryCacheEndpoint string `json:"entry_cache_endpoint"`

	// sendgrid
	Sendgrid *sendgrid.Config `json:"sendgrid"`

	// rabbit
	RabbitMQ *rabbitmq.Config
}

// ParseConfig : get configurations related to email service from common configurations
func ParseConfig(conf *config.Config) *Config {
	return &Config{
		Host:               conf.Email.Host,
		Port:               conf.Email.Port,
		Env:                string(conf.Env),
		EntryCacheEndpoint: fmt.Sprintf("%s:%s", conf.EntryCache.Host, conf.EntryCache.Port),
		Sendgrid: &sendgrid.Config{
			APIKey:      conf.Sendgrid.APIKey,
			SenderEmail: conf.Sendgrid.SenderEmail,
			SenderName:  conf.Sendgrid.SenderName,
			CCEmail:     conf.Sendgrid.CCEmail,
			CCName:      conf.Sendgrid.CCName,
		},
		RabbitMQ: &rabbitmq.Config{
			URL:      fmt.Sprintf("%s:%s", conf.RabbitMQ.Host, conf.RabbitMQ.Port),
			User:     conf.RabbitMQ.Username,
			Password: conf.RabbitMQ.Password,
		},
	}
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tEnv:\t\t\t%s\n", conf.Env)
	fmt.Printf("\t\tHost:\t\t\t%s\n", conf.Host)
	fmt.Printf("\t\tPort:\t\t\t%s\n", conf.Port)
	fmt.Printf("\t\tEntryCacheEndpoint:\t%s\n", conf.EntryCacheEndpoint)
	fmt.Printf("\t\tSendgridAPIKey:\t\t%s\n", conf.Sendgrid.APIKey)
	fmt.Printf("\t\tSendgridSenderEmail:\t%s\n", conf.Sendgrid.SenderEmail)
	fmt.Printf("\t\tSendgridSenderName:\t%s\n", conf.Sendgrid.SenderName)
	fmt.Printf("\t\tSendgridCCEmail:\t%s\n", conf.Sendgrid.CCEmail)
	fmt.Printf("\t\tSendgridCCName:\t\t%s\n", conf.Sendgrid.CCName)
	fmt.Printf("\t\tRabbitMQURL:\t\t%s\n", conf.RabbitMQ.URL)
	fmt.Printf("\t\tRabbitMQUser:\t\t%s\n", conf.RabbitMQ.User)
	fmt.Printf("\t\tRabbitMQPwd:\t\t%s\n", conf.RabbitMQ.Password)
}
