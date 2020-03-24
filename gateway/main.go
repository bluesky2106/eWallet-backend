package main

import (
	"fmt"
	"time"

	commonConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/gateway/api"
	"github.com/bluesky2106/eWallet-backend/gateway/config"
	"github.com/bluesky2106/eWallet-backend/gateway/services"
	"github.com/bluesky2106/eWallet-backend/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	conf := commonConfig.ParseConfig("config.json", "../config")
	logger = log.InitLogger(conf.Env)

	gwConf := &config.Config{
		Env:                string(conf.Env),
		Host:               conf.APIGateway.Host,
		Port:               conf.APIGateway.Port,
		EntryCacheEndpoint: fmt.Sprintf("%s:%s", conf.EntryCache.Host, conf.EntryCache.Port),
		TokenSecretKey:     conf.TokenSecretKey,
		CryptoPassphase:    conf.CryptoPassphase,
	}
	gwConf.Print()

	productSrv := services.NewProductService(gwConf)
	userSrv := services.NewUserService(gwConf)

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

	svr := api.NewServer(gwConf, router, userSrv, productSrv)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)
	if err := router.Run(fmt.Sprintf("%s:%s", gwConf.Host, gwConf.Port)); err != nil {
		logger.Error("router.Run", zap.Error(err))
	}
}
