package config

import "fmt"

// Config : gateway configuration
type Config struct {
	Host string `json: "host"`
	Port string `json: "port"`
	Env  string `json: "env"`

	// services
	EntryCacheEndpoint string `json:"entryCacheEndpoint"`
	TokenSecretKey     string `json:"tokenSecretKey"`
	CryptoPassphase    string `json:"cryptoPassphase"`
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tEnv:\t%s\n", conf.Env)
	fmt.Printf("\t\tHost:\t%s\n", conf.Host)
	fmt.Printf("\t\tPort:\t%s\n", conf.Port)
	fmt.Printf("\t\tEntryCacheEndpoint:\t%s\n", conf.EntryCacheEndpoint)
	fmt.Printf("\t\tTokenSecretKey:\t%s\n", conf.TokenSecretKey)
	fmt.Printf("\t\tCryptoPassphase:\t%s\n", conf.CryptoPassphase)
}
