package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config : configurations
type Config struct {
	Postgres
	MySQL
	Mongo

	RabbitMQ
	Redis

	APIGateway
	EntryCache
	EntryStore
	Asset
	EntryScan
	Event
	PushNotification
	Email
	Program
	BOController

	Sendgrid

	Env Environment
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
	setDefaultMySQL()
	setDefaultMongo()

	setDefaultRabbitMQ()
	setDefaultRedis()

	setDefaultAPIGateway()
	setDefaultEntryCache()
	setDefaultEntryStore()
	setDefaultAsset()
	setDefaultEntryScan()
	setDefaultEvent()
	setDefaultPushNotification()
	setDefaultEmail()
	setDefaultProgram()
	setDefaultBOController()

	setDefaultSendgrid()

	setDefaultEnvironment()
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

// Print configurations for checking
func (conf *Config) Print() {
	conf.printPostgresConfig()
	conf.printMySQLConfig()
	conf.printMongoConfig()

	conf.printRabbitMQConfig()
	conf.printRedisConfig()

	conf.printAPIGatewayConfig()
	conf.printEntryCacheConfig()
	conf.printEntryStoreConfig()
	conf.printAssetConfig()
	conf.printEntryScanConfig()
	conf.printEventConfig()
	conf.printPushNotificationConfig()
	conf.printEmailConfig()
	conf.printProgramConfig()
	conf.printBOControllerConfig()

	conf.printSendgridConfig()

	conf.printEnvironmentConfig()
}
