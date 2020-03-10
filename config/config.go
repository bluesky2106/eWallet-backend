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
	Postgres
	APIGateway
	EntryCache
	EntryStore
	Asset
	EntryScan
	Event
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
	setDefaultEntryStore()
	setDefaultAsset()
	setDefaultEntryScan()
	setDefaultEvent()
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
	conf.printPostgresConfig()
	conf.printAPIGatewayConfig()
	conf.printEntryCacheConfig()
	conf.printEntryStoreConfig()
	conf.printAssetConfig()
	conf.printEntryScanConfig()
	conf.printEventConfig()
}
