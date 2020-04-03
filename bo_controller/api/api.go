package api

import (
	"github.com/gin-gonic/gin"

	"github.com/bluesky2106/eWallet-backend/bo_controller/config"
	"github.com/bluesky2106/eWallet-backend/bo_controller/servers"
)

// Server : struct
type Server struct {
	config  *config.Config
	g       *gin.Engine
	userSrv *servers.UserSrv
}

// NewServer : userSvc, walletSvc, assetSvc, config
func NewServer(config *config.Config,
	g *gin.Engine,
	userSrv *servers.UserSrv,
) *Server {
	return &Server{
		config:  config,
		g:       g,
		userSrv: userSrv,
	}
}
