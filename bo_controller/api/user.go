package api

import (
	"net/http"

	"github.com/bluesky2106/eWallet-backend/bo_controller/serializers"
	"github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/gin-gonic/gin"
)

// Authenticate an user
func (s *Server) Authenticate(c *gin.Context) (*models.User, error) {
	// 1. Parse request body
	var req serializers.UserLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		err := errs.New(errs.ECInvalidArgument, err.Error())
		return nil, err
	}

	// 2. Authenticate
	user, err := s.userSrv.Authenticate(&req)
	if err != nil {
		return nil, errs.WithMessage(err, "s.userSrv.Authenticate")
	}

	return user, nil
}

// Register a user
func (s *Server) Register(c *gin.Context) {
	// 1. Parse request body
	var req serializers.UserRegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		err := errs.New(errs.ECInvalidArgument, err.Error())
		respondError(c, http.StatusBadRequest, err)
		return
	}

	// 2. Register new user
	user, err := s.userSrv.Register(&req)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err, "s.userSrv.Register")
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

// UpdateUserProfile : [PUT] /auth/user-profile
func (s *Server) UpdateUserProfile(c *gin.Context) {
	user, err := s.userFromContext(c)

	if err != nil {
		respondError(c, http.StatusInternalServerError, err, "s.userFromContext")
		return
	}

	var req serializers.UserUpdateProfilerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		err = errs.New(errs.ECInvalidArgument, err.Error())
		respondError(c, http.StatusBadRequest, err)
		return
	}

	user, err = s.userSrv.UpdateUserProfile(user, &req)
	if err != nil {
		respondError(c, http.StatusBadRequest, err, "s.userSrv.UpdateUserProfile")
		return
	}

	c.JSON(http.StatusOK, Resp{Result: user, Error: nil})
}

// ChangePwd : [POST] /user-change-pwd
func (s *Server) ChangePwd(c *gin.Context) {
	var req *serializers.UserChangePwdReq

	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, errs.New(errs.ECInvalidArgument, err.Error()))
		return
	}

	user, err := s.userFromContext(c)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err, "s.userFromContext")
		return
	}

	result, err := s.userSrv.ChangePwd(user, req)
	if err != nil {
		respondError(c, http.StatusInternalServerError, err, "s.userSrv.ChangePwd")
		return
	}

	c.JSON(http.StatusOK, Resp{Result: result, Error: nil})
}
