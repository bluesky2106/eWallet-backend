package services

import (
	"context"
	"strings"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/gateway/config"
	"github.com/bluesky2106/eWallet-backend/models"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

// IUserService : interface of user service
type IUserService interface {
	ReadUserEmail(email string) (*pb.ReadUserRes, error)
	CreateUser(fullName, email, pwdHashed string) (*pb.CreateUserRes, error)
}

// UserService : user service
type UserService struct {
	IUserService

	conf *config.Config
}

// NewUserService : config, rabbitmq, user message
func NewUserService(conf *config.Config) *UserService {
	return &UserService{
		conf: conf,
	}
}

// ReadUserEmail : call entry cache to read user info
//
// params: [email]
func (u *UserService) ReadUserEmail(email string) (*pb.ReadUserRes, error) {
	conn, err := grpc.Dial(u.conf.EntryCacheEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()

	c := pb.NewUserSrvClient(conn)
	r, err := c.ReadUser(context.Background(), &pb.ReadUserReq{
		Req: &pb.BaseReq{
			Action:     pb.Action_ACTION_READ,
			Message:    pb.Message_MESSAGE_READ_USER_BY_EMAIL,
			ObjectType: pb.Object_OBJECT_USER,
		},
		User: &pb.UserInfo{
			Email: email,
		},
	})

	if err != nil {
		return nil, errs.WithMessage(err, "c.ReadUser")
	}

	return r, nil
}

// CreateUser : call entry cache to create a new user
//
// params: [fullName], [email], and [password hashed]
func (u *UserService) CreateUser(fullName, email, pwdHashed string) (*pb.CreateUserRes, error) {
	conn, err := grpc.Dial(u.conf.EntryCacheEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()

	c := pb.NewUserSrvClient(conn)
	r, err := c.CreateUser(context.Background(), &pb.CreateUserReq{
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
	})

	if err != nil {
		return nil, errs.WithMessage(err, "c.CreateUser")
	}

	return r, nil
}

// Authenticate : user login request
func (u *UserService) Authenticate(req *models.UserLoginReq) (*models.User, error) {
	err := u.validateUserLoginReq(req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.validateUserLoginReq")
	}

	r, err := u.ReadUserEmail(req.Email)
	if err != nil {
		return nil, errs.WithMessage(err, "u.userMsg.ReadUserEmail")
	}

	user := &models.User{
		ID:                 uint(r.GetUser().Id),
		Email:              r.GetUser().Email,
		FullName:           r.GetUser().FullName,
		UserName:           r.GetUser().Username,
		Keystore:           r.GetUser().Keystore,
		CryptoPassphase:    u.conf.CryptoPassphase,
		EnableNotification: r.GetUser().EnableNotification,
		// UserWallets:        r.GetWallet(),
	}

	if err := bcrypt.CompareHashAndPassword([]byte(r.GetUser().Password), []byte(req.Password)); err != nil {
		err = errs.New(errs.ECSystemError, err.Error())
		return nil, errs.WithMessage(err, "bcrypt.CompareHashAndPassword")
	}

	return user, nil
}

// Register user: full name, email, password, confirm password
func (u *UserService) Register(req *models.UserRegisterReq) (*models.User, error) {
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
		return nil, errs.WithMessage(err, "u.userMsg.CreateUser")
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
