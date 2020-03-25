package main

import (
	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/entry_store/config"
	"github.com/bluesky2106/eWallet-backend/libs/mysql"
	"github.com/bluesky2106/eWallet-backend/log"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	// 1. Get global config
	conf := commonConfig.ParseConfig("config.json", "../config")
	// conf.Print()

	// 2. Init logger
	logger = log.InitLogger(conf.Env)

	// 3. Extract store config
	storeConf := &config.Config{
		Env:       string(conf.Env),
		Host:      conf.EntryCache.Host,
		Port:      conf.EntryCache.Port,
		MySQLHost: conf.MySQL.Host,
		MySQLPort: conf.MySQL.Port,
		MySQLDB:   conf.MySQL.DBName,
		MySQLUser: conf.MySQL.Username,
		MySQLPwd:  conf.MySQL.Password,
	}
	storeConf.Print()

	// 4. Init DAO
	_, err := mysql.New(&mysql.Config{
		DBName:   storeConf.MySQLDB,
		Host:     storeConf.MySQLHost,
		Port:     storeConf.MySQLPort,
		Username: storeConf.MySQLUser,
		Password: storeConf.MySQLPwd,
	}, storeConf.Env)
	if err != nil {
		logger.Error("failed to init DAO:", zap.Error(err))
	}
}
