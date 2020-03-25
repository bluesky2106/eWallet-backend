package client

import (
	"context"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"google.golang.org/grpc"
)

// UserSvc : user service client
type UserSvc struct {
	pb.UnimplementedUserSvcServer

	serverURL string
}

// NewUserService : create user service
func NewUserService(serverURL string) *UserSvc {
	return &UserSvc{
		serverURL: serverURL,
	}
}

// CreateUser : create new user
func (UserSvc *UserSvc) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.CreateUser(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.CreateUser")
	}
	return res, nil
}

// ReadUser : read user
func (UserSvc *UserSvc) ReadUser(ctx context.Context, req *pb.ReadUserReq) (*pb.ReadUserRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.ReadUser(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.ReadUser")
	}
	return res, nil
}

// UpdateUser : update user
func (UserSvc *UserSvc) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.UpdateUser(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.UpdateUser")
	}
	return res, nil
}

// DeleteUser : delete user from database
func (UserSvc *UserSvc) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.DeleteUser(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.DeleteUser")
	}
	return res, nil
}

// ChangePwd : Change user password
func (UserSvc *UserSvc) ChangePwd(ctx context.Context, req *pb.ChangePwdReq) (*pb.ChangePwdRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.ChangePwd(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.ChangePwd")
	}
	return res, nil
}

// ForgotPassword : forgot password case
func (UserSvc *UserSvc) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.ForgotPassword(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.ForgotPassword")
	}
	return res, nil
}

// CheckAuthenticationCode : check authentication code
func (UserSvc *UserSvc) CheckAuthenticationCode(ctx context.Context, req *pb.CheckAuthenticationCodeReq) (*pb.CheckAuthenticationCodeRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.CheckAuthenticationCode(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.CheckAuthenticationCode")
	}
	return res, nil
}

// ResetPassword : reset password case
func (UserSvc *UserSvc) ResetPassword(ctx context.Context, req *pb.ResetPasswordReq) (*pb.ResetPasswordRes, error) {
	conn, err := grpc.Dial(UserSvc.serverURL, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()
	client := pb.NewUserSvcClient(conn)
	res, err := client.ResetPassword(ctx, req)
	if err != nil {
		return nil, errs.WithMessage(err, "client.CheckAuthenticationCode")
	}
	return res, nil
}
