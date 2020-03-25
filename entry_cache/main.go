package main

import (
	"fmt"
	"net"

	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/entry_cache/config"
	"github.com/bluesky2106/eWallet-backend/entry_cache/servers"
	"github.com/bluesky2106/eWallet-backend/libs/redis"
	"github.com/bluesky2106/eWallet-backend/log"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var logger *zap.Logger

func main() {
	conf := commonConfig.ParseConfig("config.json", "../config")
	// conf.Print()
	logger = log.InitLogger(conf.Env)

	cacheConf := &config.Config{
		Env:                string(conf.Env),
		Host:               conf.EntryCache.Host,
		Port:               conf.EntryCache.Port,
		EntryStoreEndpoint: fmt.Sprintf("%s:%s", conf.EntryStore.Host, conf.EntryStore.Port),
		RedisHost:          conf.Redis.Host,
		RedisPort:          conf.Redis.Port,
		RedisDB:            conf.Redis.DB,
		RedisPwd:           conf.Redis.Password,
	}
	cacheConf.Print()

	redisClient := initRedis(cacheConf)

	// lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cacheConf.Host, cacheConf.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cacheConf.Port))
	if err != nil {
		logger.Error("failed to listen:", zap.Error(err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserSvcServer(grpcServer, servers.NewUserServer(cacheConf, redisClient))
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("router.Run", zap.Error(err))
	}
}

func initRedis(conf *config.Config) *redis.Client {
	redisClient, err := redis.Init(&redis.Config{
		Addr:     fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPwd,
		DB:       conf.RedisDB,
	})
	if err != nil {
		logger.Error("cannot connect redis:", zap.Error(err))
	}

	return redisClient
}
