package sendgrid

import "github.com/bluesky2106/eWallet-backend/config"

// Config : sendgrid configurations
type Config struct {
	APIKey      string `json:"apiKey"`
	SenderEmail string `json:"senderEmail"`
	SenderName  string `json:"senderName"`
	CCEmail     string `json:"ccEmail"`
	CCName      string `json:"ccName"`
}

// ParseConfig : extract sendgrid config from common config
func ParseConfig(conf *config.Config) *Config {
	return &Config{
		APIKey:      conf.Sendgrid.APIKey,
		SenderEmail: conf.Sendgrid.SenderEmail,
		SenderName:  conf.Sendgrid.SenderName,
		CCEmail:     conf.Sendgrid.CCEmail,
		CCName:      conf.Sendgrid.CCName,
	}
}
