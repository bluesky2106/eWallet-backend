package config

import "fmt"

// Config : entry-cache configuration
type Config struct {
	Host string `json: "host"`
	Port string `json: "port"`
	Env  string `json: "env"`

	// redis
	RedisDB   int    `json: "redisDB"`
	RedisPwd  string `json: "redisPwd"`
	RedisHost string `json: "redisHost"`
	RedisPort string `json: "redisPort"`

	// services
	EntryStoreEndpoint string `json:"entryStoreEndpoint"`
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tEnv:\t%s\n", conf.Env)
	fmt.Printf("\t\tHost:\t%s\n", conf.Host)
	fmt.Printf("\t\tPort:\t%s\n", conf.Port)
	fmt.Printf("\t\tRedisHost:\t%s\n", conf.RedisHost)
	fmt.Printf("\t\tRedisPort:\t%s\n", conf.RedisPort)
	fmt.Printf("\t\tRedisDB:\t%d\n", conf.RedisDB)
	fmt.Printf("\t\tRedisPwd:\t%s\n", conf.RedisPwd)
	fmt.Printf("\t\tEntryStoreEndpoint:\t%s\n", conf.EntryStoreEndpoint)
}
