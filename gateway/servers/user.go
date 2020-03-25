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

// IUserSrv : interface of user server
type IUserSrv interface {
	ReadUserByEmail(email string) (*pb.ReadUserRes, error)
	CreateUser(fullName, email, pwdHashed string) (*pb.CreateUserRes, error)
}

// UserSrv : user server
type UserSrv struct {
	IUserSrv

	userSvc *client.UserSvc

	conf *config.Config
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
func (u *UserSrv) ReadUserByEmail(email string) (*pb.ReadUserRes, error) {
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

	return u.userSvc.ReadUser(context.Background(), req)
}

// CreateUser : call entry cache to create a new user
//
// params: [fullName], [email], and [password hashed]
func (u *UserSrv) CreateUser(fullName, email, pwdHashed string) (*pb.CreateUserRes, error) {
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

	return u.userSvc.CreateUser(context.Background(), req)
}

// Authenticate : user login request
func (u *UserSrv) Authenticate(req *models.UserLoginReq) (*models.User, error) {
	err := u.validateUserLoginReq(req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.validateUserLoginReq")
	}

	r, err := u.ReadUserByEmail(req.Email)
	if err != nil {
		return nil, errs.WithMessage(err, "u.ReadUserByEmail")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(r.GetUser().Password), []byte(req.Password)); err != nil {
		err = errs.New(errs.ECSystemError, err.Error())
		return nil, errs.WithMessage(err, "bcrypt.CompareHashAndPassword")
	}

	return &models.User{
		ID:                 uint(r.GetUser().Id),
		Email:              r.GetUser().Email,
		FullName:           r.GetUser().FullName,
		UserName:           r.GetUser().Username,
		Keystore:           r.GetUser().Keystore,
		CryptoPassphase:    u.conf.CryptoPassphase,
		EnableNotification: r.GetUser().EnableNotification,
		// UserWallets:        r.GetWallet(),
	}, nil
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

	r, err := u.CreateUser(req.FullName, req.Email, string(hashed))
	if err != nil {
		return nil, errs.WithMessage(err, "u.CreateUser")
	}

	return &models.User{
		ID:                 uint(r.GetUser().Id),
		Email:              r.GetUser().Email,
		FullName:           r.GetUser().FullName,
		UserName:           r.GetUser().Username,
		Keystore:           r.GetUser().Keystore,
		CryptoPassphase:    u.conf.CryptoPassphase,
		EnableNotification: r.GetUser().EnableNotification,
		// UserWallets:        r.GetWallet(),
	}, nil
}
