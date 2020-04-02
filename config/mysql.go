package config

import (
	"github.com/spf13/viper"
)

// MySQL configurations
type MySQL struct {
	DBName   string `json: "dbName"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultMySQL(service string) {
	viper.SetDefault(service+".mysql.dbName", "test_db")
	viper.SetDefault(service+".mysql.username", "root")
	viper.SetDefault(service+".mysql.password", "")
	viper.SetDefault(service+".mysql.host", "localhost")
	viper.SetDefault(service+".mysql.port", "3306")
}
