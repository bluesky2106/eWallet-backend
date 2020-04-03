package main

import (
	"fmt"
	"time"

	"github.com/bluesky2106/eWallet-backend/bo_controller/api"
	"github.com/bluesky2106/eWallet-backend/bo_controller/config"
	"github.com/bluesky2106/eWallet-backend/bo_controller/servers"
	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func main() {
	// 1. Get global config
	conf := commonConfig.ParseConfig("config.json", "../config")

	// 2. Init logger
	logger = log.InitLogger(conf.Env)

	// 3. Extract store config
	boConf := config.ParseConfig(conf)
	boConf.Print()

	userSrv := servers.NewUserServer(boConf)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	svr := api.NewServer(boConf, router, userSrv)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)
	if err := router.Run(fmt.Sprintf("%s:%s", boConf.Host, boConf.Port)); err != nil {
		logger.Error("router.Run", zap.Error(err))
	}
}
