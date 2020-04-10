package mysql

import (
	"fmt"
)

// Config : MySQL configurations
type Config struct {
	ConnURL string
}

// ParseConfig : get configurations related to bo entry store from common configurations
func ParseConfig(user, pwd, host, port, dbName string) *Config {
	return &Config{
		ConnURL: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", user, pwd, host, port, dbName),
	}
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\t Database Connection URL:\t\t%s\n", conf.ConnURL)
}
