package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Redis configurations
type Redis struct {
	DB       int    `json: "db"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultRedis() {
	viper.SetDefault("redis.db", "rabbitmq")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "5672")
}

func (conf *Config) printRedisConfig() {
	fmt.Println("------------- Redis configurations -------------")
	fmt.Println("Redis DB is\t", conf.Redis.DB)
	fmt.Println("Redis Host is\t", conf.Redis.Host)
	fmt.Println("Redis Port is\t", conf.Redis.Port)
	fmt.Println("Redis Pass is\t", conf.Redis.Password)
}
