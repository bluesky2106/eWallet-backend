package servers

import (
	"context"
	"strings"

	"github.com/bluesky2106/eWallet-backend/bo_controller/config"
	"github.com/bluesky2106/eWallet-backend/bo_controller/serializers"
	"github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/grpc_services/client"
	"github.com/bluesky2106/eWallet-backend/libs/rabbitmq"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"golang.org/x/crypto/bcrypt"
)

// UserSrv : user server
type UserSrv struct {
	conf *config.Config
	rbmq *rabbitmq.RabbitMQ

	userSvc *client.UserSvc
}

// NewUserServer : config, rabbitmq, user message
func NewUserServer(conf *config.Config) *UserSrv {
	return &UserSrv{
		conf:    conf,
		rbmq:    rabbitmq.Init(conf.RabbitMQ, []rabbitmq.QueueName{rabbitmq.QueueEmailService}),
		userSvc: client.NewUserService(conf.BOEntryStoreEndpoint),
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

	return user, nil
}

// ReadUserByID : call entry store to read user info
//
// params: [id]
func (u *UserSrv) ReadUserByID(id uint64) (*models.User, error) {
	req := &pb.ReadUserReq{
		Req: &pb.BaseReq{
			Action:     pb.Action_ACTION_READ,
			Message:    pb.Message_MESSAGE_READ_USER_BY_ID,
			ObjectType: pb.Object_OBJECT_USER,
		},
		User: &pb.UserInfo{
			Id: id,
		},
	}

	res, err := u.userSvc.ReadUser(context.Background(), req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.userSvc.ReadUser")
	}

	user := models.ConvertPbUserToUser(res.GetUser())

	return user, nil
}

// CreateUser : call entry store to create a new user
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

	return user, nil
}

// Authenticate : user login request
func (u *UserSrv) Authenticate(req *serializers.UserLoginReq) (*models.User, error) {
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

	return user, nil
}

// Register user: full name, email, password, confirm password
func (u *UserSrv) Register(req *serializers.UserRegisterReq) (*models.User, error) {
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

	return user, nil
}

// UpdateUserProfile : [uReq : full name, email]
func (u *UserSrv) UpdateUserProfile(user *models.User, uReq *serializers.UserUpdateProfilerReq) (*models.User, error) {
	err := u.validateUserUpdateProfileReq(uReq)
	if err != nil {
		return nil, errs.WithMessage(err, "u.validateUserRegisterReq")
	}

	fullName := user.FullName
	if uReq.FullName != "" {
		fullName = uReq.FullName
	}

	email := user.Email
	if uReq.Email != "" {
		email = uReq.Email
	}

	req := &pb.UpdateUserReq{
		Req: &pb.BaseReq{
			Message:    pb.Message_MESSAGE_UPDATE_USER,
			ObjectType: pb.Object_OBJECT_USER,
			Action:     pb.Action_ACTION_UPDATE,
		},
		User: &pb.UserInfo{
			Id:       user.ID,
			FullName: fullName,
			Username: strings.Split(email, "@")[0],
			Email:    email,
			Password: user.Password,
		},
	}

	res, err := u.userSvc.UpdateUser(context.Background(), req)
	if err != nil {
		return nil, errs.WithMessage(err, "u.userSvc.UpdateUser")
	}

	newUser := models.ConvertPbUserToUser(res.GetUser())

	return newUser, nil
}

// ChangePwd : user update new password and send email to user
//
// params: [user, change pwd req]
func (u *UserSrv) ChangePwd(user *models.User, uReq *serializers.UserChangePwdReq) (bool, error) {
	// 1. Validate the request
	err := u.validateChangePwdReq(uReq)
	if err != nil {
		return false, errs.WithMessage(err, "u.validateChangePwdReq")
	}

	// 2. Check whether old password matched the one stored in DB
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(uReq.OldPassword)); err != nil {
		return false, errs.New(errs.ECChangePasswordOldPwdNotSame)
	}

	// 3. Generate new password hash
	newHashed, err := bcrypt.GenerateFromPassword([]byte(uReq.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return false, errs.BCRYPTGenerateFromPasswordError(err)
	}

	// 4. Update new password
	req := &pb.ChangePwdReq{
		Req: &pb.BaseReq{
			Message:    pb.Message_MESSAGE_CHANGE_PWD_USER,
			ObjectType: pb.Object_OBJECT_USER,
			Action:     pb.Action_ACTION_UPDATE,
		},
		User: &pb.UserInfo{
			Id:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
			Username: user.UserName,
			Password: string(newHashed),
		},
	}

	res, err := u.userSvc.ChangePwd(context.Background(), req)
	if err != nil {
		return false, errs.WithMessage(err, "u.userSvc.ChangePwd")
	}

	// 5. Send email to inform the user
	if res.GetResult() {
		emailInfo := pb.EmailInfo{
			TemplateId: res.GetTemplateId(),
			Receivers: []*pb.Receiver{
				&pb.Receiver{
					ToName:  res.GetUser().GetFullName(),
					ToEmail: res.GetUser().GetEmail(),
				},
			},
			Data: map[string]string{
				"name":      res.GetUser().GetFullName(),
				"date_time": res.GetDateTime(),
				"email":     res.GetUser().GetEmail(),
				"location":  res.GetLocation(),
			},
		}
		err := u.sendEmail(emailInfo)
		if err != nil {
			return false, errs.WithMessage(err, "send change password email failed")
		}
	}

	return res.GetResult(), nil
}
