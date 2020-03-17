package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Sendgrid configurations
type Sendgrid struct {
	APIKey      string `json:"apiKey"`
	SenderEmail string `json:"senderEmail"`
	SenderName  string `json:"senderName"`
	CCEmail     string `json:"ccEmail"`
	CCName      string `json:"ccName"`
}

func setDefaultSendgrid() {
	viper.SetDefault("sendgrid.apiKey", "")
	viper.SetDefault("sendgrid.senderEmail", "")
	viper.SetDefault("sendgrid.senderName", "")
	viper.SetDefault("sendgrid.ccEmail", "")
	viper.SetDefault("sendgrid.ccName", "")
}

func (conf *Config) printSendgridConfig() {
	fmt.Println("------------ Sendgrid configurations -----------")
	fmt.Println("Sendgrid APIKey is\t", conf.Sendgrid.APIKey)
	fmt.Println("Sendgrid SenderEmail is\t", conf.Sendgrid.SenderEmail)
	fmt.Println("Sendgrid SenderName is\t", conf.Sendgrid.SenderName)
	fmt.Println("Sendgrid CCEmail is\t", conf.Sendgrid.CCEmail)
	fmt.Println("Sendgrid CCName is\t", conf.Sendgrid.CCName)
}
