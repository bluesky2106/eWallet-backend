package services

import (
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/models"
)

func (u *UserService) validateUserRegisterReq(req *models.UserRegisterReq) error {

	if req.Email == "" {
		return errs.New(errs.ECInvalidEmail)
	}

	if req.Password == "" || req.ConfirmPassword == "" {
		return errs.New(errs.ECInvalidPassword)
	}

	if req.Password != req.ConfirmPassword {
		return errs.New(errs.ECPasswordMismatch)
	}

	return nil
}

func (u *UserService) validateUserLoginReq(req *models.UserLoginReq) error {

	if req.Email == "" {
		return errs.New(errs.ECInvalidEmail)
	}

	if req.Password == "" {
		return errs.New(errs.ECInvalidPassword)
	}

	return nil
}
