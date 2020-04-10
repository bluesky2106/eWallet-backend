package main

import (
	"fmt"
	"net"

	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/entry_store/config"
	"github.com/bluesky2106/eWallet-backend/entry_store/models"
	"github.com/bluesky2106/eWallet-backend/entry_store/servers"
	"github.com/bluesky2106/eWallet-backend/libs/mysql"
	"github.com/bluesky2106/eWallet-backend/log"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	logger *zap.Logger
	tables = []interface{}{(*models.User)(nil), (*models.ProductGroup)(nil), (*models.Unit)(nil), (*models.ProductInfo)(nil)}
)

func main() {
	// 1. Get global config
	conf := commonConfig.ParseConfig("config.json", "../config")
	// conf.Print()

	// 2. Init logger
	logger = log.InitLogger(conf.Env)

	// 3. Extract store config
	storeConf := &config.Config{
		Env:       string(conf.Env),
		Host:      conf.EntryStore.Host,
		Port:      conf.EntryStore.Port,
		MySQLHost: conf.EntryStore.MySQL.Host,
		MySQLPort: conf.EntryStore.MySQL.Port,
		MySQLDB:   conf.EntryStore.MySQL.DBName,
		MySQLUser: conf.EntryStore.MySQL.Username,
		MySQLPwd:  conf.EntryStore.MySQL.Password,
	}
	storeConf.Print()

	// 4. Init DAO
	dbConf := mysql.ParseConfig(storeConf.MySQLUser, storeConf.MySQLPwd, storeConf.MySQLHost, storeConf.MySQLPort, storeConf.MySQLDB)
	dao, err := mysql.New(dbConf, storeConf.Env)
	if err != nil {
		logger.Error("failed to init DAO:", zap.Error(err))
	}
	// 5. AutoMigrate
	err = dao.AutoMigrate(tables)
	if err != nil {
		logger.Error("failed to automigrate:", zap.Error(err))
	}
	// 6. Add foreign keys
	err = dao.AddForeignKey((*models.ProductInfo)(nil), "product_group_id", "product_groups(g_id)")
	if err != nil {
		logger.Error("failed to add foreign keys:", zap.Error(err))
	}

	// 5. Init grpc server
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
