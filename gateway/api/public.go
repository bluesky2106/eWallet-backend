package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultWelcome : ...
func (s *Server) DefaultWelcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Homepage")
}

// Welcome : ...
func (s *Server) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, Resp{Result: "REST API Homepage"})
}
