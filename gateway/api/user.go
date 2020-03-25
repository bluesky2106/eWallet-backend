package api

import (
	"net/http"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/gin-gonic/gin"

	"github.com/bluesky2106/eWallet-backend/models"
)

// Authenticate an user
func (s *Server) Authenticate(c *gin.Context) (*models.User, error) {
	// 1. Parse request body
	var req models.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		err := errs.New(errs.ECInvalidArgument, err.Error())
		return nil, err
	}

	// 2. Authenticate
	user, err := s.userSrv.Authenticate(&req)
	if err != nil {
		return nil, errs.WithMessage(err, "s.userSvc.Authenticate")
	}

	return user, nil
}

// Register a user
func (s *Server) Register(c *gin.Context) {
	// 1. Parse request body
	var req models.UserRegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		err := errs.New(errs.ECInvalidArgument, err.Error())
		respondError(c, http.StatusBadRequest, err)
		return
	}

	// 2. Register new user
	user, err := s.userSrv.Register(&req)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err, "s.userSvc.Register")
		return
	}

	// 3. Write to html
	c.JSON(http.StatusOK, Resp{Result: user, Error: nil})
}

// UserProfile : [GET] /auth/user-profile
func (s *Server) UserProfile(c *gin.Context) {
	user, err := s.userFromContext(c)

	if err != nil {
		respondError(c, http.StatusInternalServerError, err, "s.userFromContext")
		return
	}

	c.JSON(http.StatusOK, Resp{Result: user, Error: nil})
}
