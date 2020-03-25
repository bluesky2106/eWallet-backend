package api

import (
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) userFromContext(c *gin.Context) (*models.User, error) {
	userIDVal, ok := c.Get(userIDKey)
	if !ok {
		return nil, errs.New(errs.ECInvalidCredentials, "failed to get userIDKey from context")
	}

	userID := userIDVal.(float64)
	user, err := s.userSrv.ReadUserByID(uint(userID))

	if err != nil {
		return nil, errs.WithMessage(err, "s.userSvc.FindByID")
	}

	return user, nil
}
