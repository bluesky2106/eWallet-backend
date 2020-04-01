package api

import (
	"net/http"

	"github.com/bluesky2106/eWallet-backend/entry_store/models"
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/gin-gonic/gin"
)

// AddProductGroup : add a new product group
func (s *Server) AddProductGroup(c *gin.Context) {
	var req models.ProductGroup
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusInternalServerError, errs.New(errs.ECInvalidArgument, err.Error()))
		return
	}

	result, err := s.productSrv.AddProductGroup(&req)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err, "s.assetSvc.AddNewTnx")
		return
	}

	c.JSON(http.StatusOK, Resp{Result: result, Error: nil})
}
