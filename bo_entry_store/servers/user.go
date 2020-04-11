package servers

import (
	"context"

	"github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/libs/mysql"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
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
	// 1. Validate request
	if !u.isValidUserRequest(req.GetReq()) {
		return nil, errs.New(errs.ECInvalidMessage)
	}

	// 2. Check if user with that email exists
	email := req.GetUser().GetEmail()
	user, _ := u.dao.FindOneByQuery(models.User{}, map[string]interface{}{
		"email = ?": email,
	})
	if user != nil {
		return nil, errs.New(errs.ECEmailAlreadyExists)
	}

	// 3. Create new User
	newUser := models.ConvertPbUserToUser(req.GetUser())
	err := u.dao.WithTransaction(func() error {
		err := u.dao.Create(newUser)
		if err != nil {
			return errs.WithMessage(err, "u.dao.Create")
		}

		return nil
	})
	if err != nil {
		return nil, errs.WithMessage(err, "u.dao.WithTransaction")
	}

	// 4. Finally, returns pb.CreateUserRes
	return &pb.CreateUserRes{
		Result: true,
		User:   models.ConvertUserToPbUser(newUser),
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
		User:   models.ConvertUserToPbUser(user),
	}, nil
}

// UpdateUser : update user
func (u *UserSrv) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	// 1. Validate request
	if !u.isValidUserRequest(req.GetReq()) {
		return nil, errs.New(errs.ECInvalidMessage)
	}

	// 2. Update user
	user := models.ConvertPbUserToUser(req.GetUser())
	err := u.dao.WithTransaction(func() error {
		err := u.dao.Update(user)
		if err != nil {
			return errs.WithMessage(err, "u.dao.Update")
		}

		return nil
	})
	if err != nil {
		return nil, errs.WithMessage(err, "u.dao.WithTransaction")
	}

	return &pb.UpdateUserRes{
		Result: true,
		User:   req.GetUser(),
	}, nil
}
