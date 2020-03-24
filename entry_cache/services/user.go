package services

import (
	"context"

	"github.com/bluesky2106/eWallet-backend/entry_cache/config"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User service
type User struct {
	pb.UnimplementedUserSrvServer

	conf *config.Config
}

// NewUserService : create user service
func NewUserService(conf *config.Config) *User {
	return &User{
		conf: conf,
	}
}

// CreateUser : create new user
func (user *User) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// ReadUser : read user
func (user *User) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUser not implemented")
}

// UpdateUser : update user
func (user *User) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

// DeleteUser : delete user from database
func (user *User) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

// ChangePwd : Change user password
func (user *User) ChangePwd(ctx context.Context, req *pb.ChangePwdReq) (*pb.ChangePwdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePwd not implemented")
}

// ForgotPassword : forgot password case
func (user *User) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}

// CheckAuthenticationCode : check authentication code
func (user *User) CheckAuthenticationCode(ctx context.Context, req *pb.CheckAuthenticationCodeReq) (*pb.CheckAuthenticationCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthenticationCode not implemented")
}

// ResetPassword : reset password case
func (user *User) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
