package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// DatabaseDriver : database driver
type DatabaseDriver string

const (
	// MySQL : MySQL database
	MySQL DatabaseDriver = "mysql"
	// Postgres : Postgres database
	Postgres DatabaseDriver = "postgres"
	// Mongo : Mongo database
	Mongo DatabaseDriver = "mongo"
)

// Config : configurations
type Config struct {
	Database
}

// Database configurations exported
type Database struct {
	Name     string `json: "name"`
	Driver   string `json: "driver"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

// ParseConfig : parse configurations from global env and json file
func ParseConfig(file, path string) *Config {
	// Set default variables
	setDefaultVariables()

	// Enable VIPER to read Environment Variables
	readEnvironmentVariables()

	// Parse configurations from JSON file
	readJSONFile(file, path)

	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}

	return &conf
}

func setDefaultVariables() {
	// Set default db variables
	viper.SetDefault("database.name", "test_db")
	viper.SetDefault("database.driver", string(Postgres))
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.pass", "")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")

	// Set default server variables
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "3000")
}

func readEnvironmentVariables() {
	viper.AutomaticEnv()
}

func readJSONFile(file, path string) {
	fileName, fileType := getFileNameAndType(file)

	// Set the file name & type of the configurations file
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)

	// Set the path to look for the configurations file
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}
}

func getFileNameAndType(file string) (fileName string, fileType string) {
	strs := strings.Split(file, ".")

	if len(strs) != 2 {
		fmt.Println("Config file name must follow format of xxx.yyy")
		return "", ""
	}

	fileName = strs[0]
	fileType = strs[1]

	return fileName, fileType
}

// Print configurations for checking
func (conf *Config) Print() {
	fmt.Println("---------- Database configurations ----------")
	fmt.Println("Database name is\t", conf.Database.Name)
	fmt.Println("Database type is\t", conf.Database.Driver)
	fmt.Println("Database User is\t", conf.Database.Username)
	fmt.Println("Database Pass is\t", conf.Database.Password)
	fmt.Println("Database Host is\t", conf.Database.Host)
	fmt.Println("Database Port is\t", conf.Database.Port)

	fmt.Println("----------- Server configurations -----------")
	// fmt.Println("Server host is\t", conf.Server.Host)
	// fmt.Println("Server port is\t", conf.Server.Port)
}

// GetDatabaseConfigurations returns configurations for database
func (conf *Config) GetDatabaseConfigurations() Database {
	return conf.Database
}
