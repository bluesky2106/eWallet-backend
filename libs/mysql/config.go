package mysql

// Config : MySQL configurations
type Config struct {
	DBName   string `json: "dbName"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}
