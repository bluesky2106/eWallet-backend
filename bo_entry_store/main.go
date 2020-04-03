package main

import (
	"fmt"
	"net"

	"github.com/bluesky2106/eWallet-backend/bo_entry_store/config"
	"github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
	"github.com/bluesky2106/eWallet-backend/bo_entry_store/servers"
	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/libs/mysql"
	"github.com/bluesky2106/eWallet-backend/log"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	logger *zap.Logger
	tables = []interface{}{(*models.User)(nil)}
)

func main() {
	// 1. Get global config
	conf := commonConfig.ParseConfig("config.json", "../config")

	// 2. Init logger
	logger = log.InitLogger(conf.Env)

	// 3. Extract store config
	storeConf := &config.Config{
		Env:       string(conf.Env),
		Host:      conf.BOEntryStore.Host,
		Port:      conf.BOEntryStore.Port,
		MySQLHost: conf.BOEntryStore.MySQL.Host,
		MySQLPort: conf.BOEntryStore.MySQL.Port,
		MySQLDB:   conf.BOEntryStore.MySQL.DBName,
		MySQLUser: conf.BOEntryStore.MySQL.Username,
		MySQLPwd:  conf.BOEntryStore.MySQL.Password,
	}
	storeConf.Print()

	// 4. Init DAO
	dao, err := mysql.New(&mysql.Config{
		DBName:   storeConf.MySQLDB,
		Host:     storeConf.MySQLHost,
		Port:     storeConf.MySQLPort,
		Username: storeConf.MySQLUser,
		Password: storeConf.MySQLPwd,
	}, storeConf.Env)
	if err != nil {
		logger.Error("failed to init DAO:", zap.Error(err))
	}
	// 5. AutoMigrate
	err = dao.AutoMigrate(tables)
	if err != nil {
		logger.Error("failed to automigrate:", zap.Error(err))
	}

	// 6. Init grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", storeConf.Host, storeConf.Port))
	if err != nil {
		logger.Error("failed to listen", zap.Error(err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserSvcServer(grpcServer, servers.NewUserServer(
		dao,
	))
	grpcServer.Serve(lis)
}
