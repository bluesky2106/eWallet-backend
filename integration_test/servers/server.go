package servers

import (
	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/integration_test/config"
)

// TestConfig : test configurations
type TestConfig struct {
	Conf *config.Config
}

// TestData : test data
type TestData struct {
}

var (
	// testConf :
	testConf *TestConfig
)

// GetTestingConfig : get test configurations
func GetTestingConfig() *TestConfig {
	return testConf
}

func init() {
	testConf = new(TestConfig)

	conf := commonConfig.ParseConfig("config.json", "../config")
	testConf.Conf = config.ParseConfig(conf)
	testConf.Conf.Print()
}
