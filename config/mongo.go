package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Mongo configurations
type Mongo struct {
	DBName   string `json: "dbName"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultMongo() {
	viper.SetDefault("mongo.dbName", "test_db")
	viper.SetDefault("mongo.username", "")
	viper.SetDefault("mongo.password", "")
	viper.SetDefault("mongo.host", "localhost")
	viper.SetDefault("mongo.port", "27017")
}

func (conf *Config) printMongoConfig() {
	fmt.Println("----------- Mongo DB configurations ------------")
	fmt.Println("Mongo DB name is\t", conf.Mongo.DBName)
	fmt.Println("Mongo User is\t\t", conf.Mongo.Username)
	fmt.Println("Mongo Pass is\t\t", conf.Mongo.Password)
	fmt.Println("Mongo Host is\t\t", conf.Mongo.Host)
	fmt.Println("Mongo Port is\t\t", conf.Mongo.Port)
}
