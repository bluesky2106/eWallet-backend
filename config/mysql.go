package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// MySQL configurations
type MySQL struct {
	DBName   string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultMySQL() {
	viper.SetDefault("mysql.dbName", "test_db")
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "")
	viper.SetDefault("mysql.host", "localhost")
	viper.SetDefault("mysql.port", "3306")
}

func (conf *Config) printMySQLConfig() {
	fmt.Println("------------ MySQL DB configurations -----------")
	fmt.Println("MySQL DB name is\t", conf.MySQL.DBName)
	fmt.Println("MySQL User is\t\t", conf.MySQL.Username)
	fmt.Println("MySQL Pass is\t\t", conf.MySQL.Password)
	fmt.Println("MySQL Host is\t\t", conf.MySQL.Host)
	fmt.Println("MySQL Port is\t\t", conf.MySQL.Port)
}
