package api

import (
	"github.com/gin-gonic/gin"

	"github.com/bluesky2106/eWallet-backend/gateway/config"
	"github.com/bluesky2106/eWallet-backend/gateway/services"
)

// Server : struct
type Server struct {
	config     *config.Config
	g          *gin.Engine
	productSvc *services.ProductService
	userSvc    *services.UserService
}

// NewServer : userSvc, walletSvc, assetSvc, config
func NewServer(config *config.Config,
	g *gin.Engine,
	userSvc *services.UserService,
	productSvc *services.ProductService,
) *Server {
	return &Server{
		config:     config,
		g:          g,
		userSvc:    userSvc,
		productSvc: productSvc,
	}
}
