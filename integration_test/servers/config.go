package servers

import (
	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/integration_test/config"
	"github.com/bluesky2106/eWallet-backend/libs/mysql"
)

var (
	testConf *TestConfig
)

// TestConfig : test configurations
type TestConfig struct {
	Conf *config.Config

	DAO   *mysql.DAO
	DAOBO *mysql.DAO
}

// GetTestingConfig : get test configurations
func GetTestingConfig() *TestConfig {
	return testConf
}

func initTestConfig() {
	testConf = new(TestConfig)

	conf := commonConfig.ParseConfig("config.json", "../../config")
	testConf.Conf = config.ParseConfig(conf)
	testConf.Conf.Print()
}
