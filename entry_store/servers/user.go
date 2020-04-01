package servers

import (
	"context"

	"github.com/bluesky2106/eWallet-backend/entry_store/models"
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/libs/mysql"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserSrv : user server
type UserSrv struct {
	pb.UnimplementedUserSvcServer

	dao *mysql.DAO
}

// NewUserServer : new user server
func NewUserServer(dao *mysql.DAO) *UserSrv {
	return &UserSrv{
		dao: dao,
	}
}

// CreateUser : create new user
func (u *UserSrv) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	// 1. Check if user with that email exists
	email := req.GetUser().GetEmail()
	user, err := u.dao.FindOneByQuery(models.User{}, map[string]interface{}{
		"email = ?": email,
	})
	if user != nil {
		return nil, errs.New(errs.ECEmailAlreadyExists)
	}

	// 2. Create new User
	newUser := &models.User{
		FullName: req.GetUser().GetFullName(),
		UserName: req.GetUser().GetUsername(),
		Email:    req.GetUser().GetEmail(),
		Password: req.GetUser().GetPassword(),
	}
	newUser.Keystore = u.generateUserKeystore()

	err = u.dao.WithTransaction(func() error {
		err := u.dao.Create(newUser)
		if err != nil {
			return errs.WithMessage(err, "u.dao.Create")
		}

		return nil
	})

	if err != nil {
		return nil, errs.WithMessage(err, "u.dao.WithTransaction")
	}

	return &pb.CreateUserRes{
		Result: true,
		User: &pb.UserInfo{
			Id:                 uint32(newUser.ID),
			FullName:           newUser.FullName,
			Email:              newUser.Email,
			Username:           newUser.UserName,
			Password:           newUser.Password,
			Keystore:           newUser.Keystore,
			EnableNotification: newUser.EnableNotification,
		},
	}, nil
}

// ReadUser : read user
func (u *UserSrv) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {
	// 1. Validate request
	if !u.isValidUserRequest(req.GetReq()) {
		return nil, errs.New(errs.ECInvalidMessage)
	}

	// 2. Find user by ID or Email
	user := new(models.User)

	switch req.GetReq().GetMessage() {
	case pb.Message_MESSAGE_READ_USER_BY_ID:
		userID := req.GetUser().GetId()
		mod, err := u.dao.FindOneByQuery(user, map[string]interface{}{
			"id = ?": userID,
		})
		if err != nil {
			return nil, errs.WithMessage(err, "u.dao.FindOneByQuery")
		}
		user = mod.(*models.User)
	case pb.Message_MESSAGE_READ_USER_BY_EMAIL:
		email := req.GetUser().GetEmail()
		mod, err := u.dao.FindOneByQuery(user, map[string]interface{}{
			"email = ?": email,
		})
		if err != nil {
			return nil, errs.WithMessage(err, "u.dao.FindOneByQuery")
		}
		user = mod.(*models.User)
	}

	return &pb.ReadUserRes{
		Result: true,
		User: &pb.UserInfo{
			Id:                 uint32(user.ID),
			FullName:           user.FullName,
			Email:              user.Email,
			Username:           user.UserName,
			Password:           user.Password,
			Keystore:           user.Keystore,
			EnableNotification: user.EnableNotification,
		},
	}, nil
}

// UpdateUser : update user
func (u *UserSrv) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// DeleteUser : delete user from database
func (u *UserSrv) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// ChangePwd : Change user password
func (u *UserSrv) ChangePwd(ctx context.Context, req *pb.ChangePwdReq) (*pb.ChangePwdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// ForgotPassword : forgot password case
func (u *UserSrv) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// CheckAuthenticationCode : check authentication code
func (u *UserSrv) CheckAuthenticationCode(ctx context.Context, req *pb.CheckAuthenticationCodeReq) (*pb.CheckAuthenticationCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// ResetPassword : reset password case
func (u *UserSrv) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
