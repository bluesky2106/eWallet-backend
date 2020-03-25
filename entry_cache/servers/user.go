package servers

import (
	"context"

	"github.com/bluesky2106/eWallet-backend/entry_cache/config"
	"github.com/bluesky2106/eWallet-backend/grpc_services/client"
	"github.com/bluesky2106/eWallet-backend/libs/redis"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

// UserSrv : user server
type UserSrv struct {
	userSvc     *client.UserSvc
	redisClient *redis.Client
}

// NewUserServer : new user server
func NewUserServer(conf *config.Config, redisClient *redis.Client) *UserSrv {
	return &UserSrv{
		userSvc:     client.NewUserService(conf.EntryStoreEndpoint),
		redisClient: redisClient,
	}
}

// CreateUser : create new user
func (u *UserSrv) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	return u.userSvc.CreateUser(ctx, req)
}

// ReadUser : read user
func (u *UserSrv) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {
	return u.userSvc.ReadUser(ctx, req)
}

// UpdateUser : update user
func (u *UserSrv) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return u.userSvc.UpdateUser(ctx, req)
}

// DeleteUser : delete user from database
func (u *UserSrv) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	return u.userSvc.DeleteUser(ctx, req)
}

// ChangePwd : Change user password
func (u *UserSrv) ChangePwd(ctx context.Context, req *pb.ChangePwdReq) (*pb.ChangePwdRes, error) {
	return u.userSvc.ChangePwd(ctx, req)
}

// ForgotPassword : forgot password case
func (u *UserSrv) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	return u.userSvc.ForgotPassword(ctx, req)
}

// CheckAuthenticationCode : check authentication code
func (u *UserSrv) CheckAuthenticationCode(ctx context.Context, req *pb.CheckAuthenticationCodeReq) (*pb.CheckAuthenticationCodeRes, error) {
	return u.userSvc.CheckAuthenticationCode(ctx, req)
}

// ResetPassword : reset password case
func (u *UserSrv) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	return u.userSvc.ResetPassword(ctx, req)
}
