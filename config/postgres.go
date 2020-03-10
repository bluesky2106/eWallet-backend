package config

import "github.com/spf13/viper"

// Postgres configurations
type Postgres struct {
	Name     string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultPostgres() {
	viper.SetDefault("postgres.name", "test_db")
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "")
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
}
