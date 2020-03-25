package api

import (
	"github.com/gin-gonic/gin"

	"github.com/bluesky2106/eWallet-backend/gateway/config"
	"github.com/bluesky2106/eWallet-backend/gateway/servers"
)

// Server : struct
type Server struct {
	config     *config.Config
	g          *gin.Engine
	productSrv *servers.ProductSrv
	userSrv    *servers.UserSrv
}

// NewServer : userSvc, walletSvc, assetSvc, config
func NewServer(config *config.Config,
	g *gin.Engine,
	userSrv *servers.UserSrv,
	productSrv *servers.ProductSrv,
) *Server {
	return &Server{
		config:     config,
		g:          g,
		userSrv:    userSrv,
		productSrv: productSrv,
	}
}
