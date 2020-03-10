package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// DatabaseDriver : database driver
type DatabaseDriver string

// Config : configurations
type Config struct {
	Postgres `json: "postgres"`

	APIGateway
	EntryCache
	EntryStore
}

// Postgres configurations
type Postgres struct {
	Name     string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

// APIGateway configurations
type APIGateway struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

// EntryCache configurations
type EntryCache struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

// EntryStore configurations
type EntryStore struct {
	Host string `json: "host"`
	Port string `json: "port"`
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
	setDefaultPostgres()
	setDefaultAPIGateway()
	setDefaultEntryCache()
}

func setDefaultPostgres() {
	viper.SetDefault("postgres.name", "test_db")
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "")
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
}

func setDefaultAPIGateway() {
	viper.SetDefault("apiGateway.host", "localhost")
	viper.SetDefault("apiGateway.port", "1")
}

func setDefaultEntryCache() {
	viper.SetDefault("entryCache.host", "localhost")
	viper.SetDefault("entryCache.port", "2")
}

func setDefaultEntryStore() {
	viper.SetDefault("entryStore.host", "localhost")
	viper.SetDefault("entryStore.port", "3")
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
	fmt.Println("Postgres DB name is\t", conf.Postgres.Name)
	fmt.Println("Postgres User is\t", conf.Postgres.Username)
	fmt.Println("Postgres Pass is\t", conf.Postgres.Password)
	fmt.Println("Postgres Host is\t", conf.Postgres.Host)
	fmt.Println("Database Port is\t", conf.Postgres.Port)

	fmt.Println("-------- API Gateway configurations ---------")
	fmt.Println("Server host is\t", conf.APIGateway.Host)
	fmt.Println("Server port is\t", conf.APIGateway.Port)

	fmt.Println("-------- Entry Cache configurations ---------")
	fmt.Println("Server host is\t", conf.EntryCache.Host)
	fmt.Println("Server port is\t", conf.EntryCache.Port)
}
