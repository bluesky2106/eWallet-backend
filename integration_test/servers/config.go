package servers

import (
	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/integration_test/config"
	"github.com/bluesky2106/eWallet-backend/libs/mysql"
)

// TestServer : test server includes conf, daos and ser
type TestServer struct {
	Conf *config.Config

	DAO   *mysql.DAO
	DAOBO *mysql.DAO
}

// NewTestServer : create new test server
func NewTestServer() *TestServer {
	server := new(TestServer)
	server.init()

	return server
}

func (ts *TestServer) init() {
	ts.initConfig()
	ts.initDAOs()
}

func (ts *TestServer) initConfig() {
	conf := commonConfig.ParseConfig("config.json", "../../config")
	ts.Conf = config.ParseConfig(conf)
	ts.Conf.Print()
}

func (ts *TestServer) initDAOs() {
	// DAO
	dao, err := mysql.New(&mysql.Config{
		ConnURL: ts.Conf.DBConn,
	}, ts.Conf.Env)
	if err != nil {
		return
	}
	ts.DAO = dao

	// BO DAO
	dao, err = mysql.New(&mysql.Config{
		ConnURL: ts.Conf.DBBOConn,
	}, ts.Conf.Env)
	if err != nil {
		return
	}
	ts.DAOBO = dao
}
