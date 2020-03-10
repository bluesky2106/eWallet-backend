package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// PushNotification configurations
type PushNotification struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

func setDefaultPushNotification() {
	viper.SetDefault("pushNotification.host", "localhost")
	viper.SetDefault("pushNotification.port", "7")
}

func (conf *Config) printPushNotificationConfig() {
	fmt.Println("------ Push Notification configurations --------")

	fmt.Println("Server host is\t", conf.PushNotification.Host)
	fmt.Println("Server port is\t", conf.PushNotification.Port)
}
