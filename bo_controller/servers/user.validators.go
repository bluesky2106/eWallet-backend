package servers

import (
	"regexp"

	"github.com/bluesky2106/eWallet-backend/bo_controller/serializers"
	errs "github.com/bluesky2106/eWallet-backend/errors"
)

// isValidEmail : email
func isValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func (u *UserSrv) validateUserRegisterReq(req *serializers.UserRegisterReq) error {
	if !isValidEmail(req.Email) {
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

func (u *UserSrv) validateUserLoginReq(req *serializers.UserLoginReq) error {

	if req.Email == "" {
		return errs.New(errs.ECInvalidEmail)
	}

	if req.Password == "" {
		return errs.New(errs.ECInvalidPassword)
	}

	return nil
}

func (u *UserSrv) validateUserUpdateProfileReq(req *serializers.UserUpdateProfilerReq) error {
	if req.FullName == "" {
		if !isValidEmail(req.Email) {
			return errs.New(errs.ECInvalidEmail)
		}
		return nil
	}

	if req.Email != "" && !isValidEmail(req.Email) {
		return errs.New(errs.ECInvalidEmail)
	}

	return nil
}

func (u *UserSrv) validateChangePwdReq(req *serializers.UserChangePwdReq) error {

	if req.NewPassword == "" || req.ConfirmNewPassword == "" || req.OldPassword == "" {
		return errs.New(errs.ECInvalidPassword)
	}

	if req.OldPassword == req.NewPassword {
		return errs.New(errs.ECChangePasswordSame)
	}

	if req.NewPassword != req.ConfirmNewPassword {
		return errs.New(errs.ECPasswordMismatch)
	}

	return nil
}
