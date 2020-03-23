package config

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
