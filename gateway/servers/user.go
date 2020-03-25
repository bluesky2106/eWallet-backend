package servers

import (
	"context"
	"strings"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/gateway/config"
	"github.com/bluesky2106/eWallet-backend/grpc_services/client"
	"github.com/bluesky2106/eWallet-backend/models"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"golang.org/x/crypto/bcrypt"
)

// UserSrv : user server
type UserSrv struct {
	conf *config.Config

	userSvc *client.UserSvc
}

// NewUserServer : config, rabbitmq, user message
func NewUserServer(conf *config.Config) *UserSrv {
	return &UserSrv{
		conf:    conf,
		userSvc: client.NewUserService(conf.EntryCacheEndpoint),
	}
}

// ReadUserByEmail : call entry cache to read user info
//
// params: [email]
func (u *UserSrv) ReadUserByEmail(email string) (*models.User, error) {
	req := &pb.ReadUserReq{
		Req: &pb.BaseReq{
			Action:     pb.Action_ACTION_READ,
			Message:    pb.Message_MESSAGE_READ_USER_BY_EMAIL,
			ObjectType: pb.Object_OBJECT_USER,
		},
		User: &pb.UserInfo{
			Email: email,
		},
	}

	res, err := u.userSvc.ReadUser(context.Background(), req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.userSvc.ReadUser")
	}

	user := models.ConvertPbUserToUser(res.GetUser())
	user.CryptoPassphase = u.conf.CryptoPassphase

	return user, nil
}

// ReadUserByID : call entry cache to read user info
//
// params: [id]
func (u *UserSrv) ReadUserByID(id uint) (*models.User, error) {
	req := &pb.ReadUserReq{
		Req: &pb.BaseReq{
			Action:     pb.Action_ACTION_READ,
			Message:    pb.Message_MESSAGE_READ_USER_BY_ID,
			ObjectType: pb.Object_OBJECT_USER,
		},
		User: &pb.UserInfo{
			Id: uint32(id),
		},
	}

	res, err := u.userSvc.ReadUser(context.Background(), req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.userSvc.ReadUser")
	}

	user := models.ConvertPbUserToUser(res.GetUser())
	user.CryptoPassphase = u.conf.CryptoPassphase

	return user, nil
}

// CreateUser : call entry cache to create a new user
//
// params: [fullName], [email], and [password hashed]
func (u *UserSrv) CreateUser(fullName, email, pwdHashed string) (*models.User, error) {
	req := &pb.CreateUserReq{
		Req: &pb.BaseReq{
			Message:    pb.Message_MESSAGE_CREATE_USER,
			ObjectType: pb.Object_OBJECT_USER,
			Action:     pb.Action_ACTION_CREATE,
		},
		User: &pb.UserInfo{
			FullName: fullName,
			Username: strings.Split(email, "@")[0],
			Email:    email,
			Password: string(pwdHashed),
		},
	}

	res, err := u.userSvc.CreateUser(context.Background(), req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.userSvc.CreateUser")
	}

	user := models.ConvertPbUserToUser(res.GetUser())
	user.CryptoPassphase = u.conf.CryptoPassphase

	return user, nil
}

// Authenticate : user login request
func (u *UserSrv) Authenticate(req *models.UserLoginReq) (*models.User, error) {
	err := u.validateUserLoginReq(req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.validateUserLoginReq")
	}

	user, err := u.ReadUserByEmail(req.Email)
	if err != nil {
		return nil, errs.WithMessage(err, "u.ReadUserByEmail")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		err = errs.New(errs.ECSystemError, err.Error())
		return nil, errs.WithMessage(err, "bcrypt.CompareHashAndPassword")
	}

	user.CryptoPassphase = u.conf.CryptoPassphase

	return user, nil
}

// Register user: full name, email, password, confirm password
func (u *UserSrv) Register(req *models.UserRegisterReq) (*models.User, error) {
	err := u.validateUserRegisterReq(req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.validateUserRegisterReq")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errs.New(errs.ECSystemError, err.Error())
		return nil, errs.WithMessage(err, "bcrypt.CompareHashAndPassword")
	}

	user, err := u.CreateUser(req.FullName, req.Email, string(hashed))
	if err != nil {
		return nil, errs.WithMessage(err, "u.CreateUser")
	}

	user.CryptoPassphase = u.conf.CryptoPassphase

	return user, nil
}
