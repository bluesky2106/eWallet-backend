package api

import (
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/gin-gonic/gin"
)

// Resp : struct
type Resp struct {
	Result interface{} `json:"Result"`
	Error  interface{} `json:"Error"`
}

func respondError(c *gin.Context, status int, err error, wrapMsg ...string) {
	e := errs.WithMessage(err, wrapMsg...)
	c.JSON(status, Resp{Error: errs.FromError(e)})
}
