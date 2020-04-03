package models

import (
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

// User : BO User struct
type User struct {
	ID       uint64
	FullName string
	UserName string
	Email    string `gorm:"unique_index"`
	Password string
	Level    uint32
	Role     uint32
}

// ConvertUserToPbUser : convert user to pb.User
func ConvertUserToPbUser(user *User) *pb.UserInfo {
	return &pb.UserInfo{
		Id:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
		Username: user.UserName,
		Level:    user.Level,
		Role:     user.Role,
	}
}

// ConvertPbUserToUser : convert pb.User to User
func ConvertPbUserToUser(user *pb.UserInfo) *User {
	return &User{
		ID:       user.GetId(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		FullName: user.GetFullName(),
		UserName: user.GetUsername(),
		Level:    user.GetLevel(),
		Role:     user.GetRole(),
	}
}
