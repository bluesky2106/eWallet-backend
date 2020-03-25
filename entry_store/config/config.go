package config

import "fmt"

// Config : entry store configurations
type Config struct {
	Host      string `json: "host"`
	Port      string `json: "port"`
	Env       string `json: "env"`
	MySQLHost string `json: "mysqlHost"`
	MySQLPort string `json: "mysqlPort"`
	MySQLDB   string `json: "mysqlDB"`
	MySQLUser string `json: "mysqlUser"`
	MySQLPwd  string `json: "mysqlPwd"`
}

// Print configurations
func (conf *Config) Print() {
	fmt.Printf("\t\tEnv:\t\t\t%s\n", conf.Env)
	fmt.Printf("\t\tHost:\t\t\t%s\n", conf.Host)
	fmt.Printf("\t\tPort:\t\t\t%s\n", conf.Port)
	fmt.Printf("\t\tMySQLHost:\t\t%s\n", conf.MySQLHost)
	fmt.Printf("\t\tMySQLPort:\t\t%s\n", conf.MySQLPort)
	fmt.Printf("\t\tMySQLDB:\t\t%s\n", conf.MySQLDB)
	fmt.Printf("\t\tMySQLUser:\t\t%s\n", conf.MySQLUser)
	fmt.Printf("\t\tMySQLPwd:\t\t%s\n", conf.MySQLPwd)
}
