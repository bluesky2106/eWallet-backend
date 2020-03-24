package servers

import (
	"context"

	"github.com/bluesky2106/eWallet-backend/entry_cache/config"
	"github.com/bluesky2106/eWallet-backend/entry_cache/services"
	"github.com/bluesky2106/eWallet-backend/libs/redis"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

// User server
type User struct {
	pb.UnimplementedUserSrvServer

	userSrv *services.User
}

// NewUserServer : new user server
func NewUserServer(conf *config.Config, redisClient *redis.Client) *User {
	return &User{
		userSrv: services.NewUserService(conf),
	}
}

// CreateUser : create new user
func (u *User) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	return u.userSrv.CreateUser(ctx, req)
}

// ReadUser : read user
func (u *User) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {
	return u.userSrv.ReadUser(ctx, req)
}

// UpdateUser : update user
func (u *User) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return u.userSrv.UpdateUser(ctx, req)
}

// DeleteUser : delete user from database
func (u *User) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	return u.userSrv.DeleteUser(ctx, req)
}

// ChangePwd : Change user password
func (u *User) ChangePwd(ctx context.Context, req *pb.ChangePwdReq) (*pb.ChangePwdRes, error) {
	return u.userSrv.ChangePwd(ctx, req)
}

// ForgotPassword : forgot password case
func (u *User) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	return u.userSrv.ForgotPassword(ctx, req)
}

// CheckAuthenticationCode : check authentication code
func (u *User) CheckAuthenticationCode(ctx context.Context, req *pb.CheckAuthenticationCodeReq) (*pb.CheckAuthenticationCodeRes, error) {
	return u.userSrv.CheckAuthenticationCode(ctx, req)
}

// ResetPassword : reset password case
func (u *User) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	return u.userSrv.ResetPassword(ctx, req)
}
