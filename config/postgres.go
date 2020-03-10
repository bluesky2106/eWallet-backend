package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Postgres configurations
type Postgres struct {
	DBName   string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultPostgres() {
	viper.SetDefault("postgres.dbName", "test_db")
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "")
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
}

func (conf *Config) printPostgresConfig() {
	fmt.Println("---------- Postgres DB configurations ----------")
	fmt.Println("Postgres DB name is\t", conf.Postgres.DBName)
	fmt.Println("Postgres User is\t", conf.Postgres.Username)
	fmt.Println("Postgres Pass is\t", conf.Postgres.Password)
	fmt.Println("Postgres Host is\t", conf.Postgres.Host)
	fmt.Println("Database Port is\t", conf.Postgres.Port)
}
