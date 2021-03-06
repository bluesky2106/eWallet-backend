package models

import (
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

// User : struct
type User struct {
	ID                 uint64
	FullName           string
	UserName           string
	Email              string `gorm:"unique_index"`
	Password           string
	Keystore           string
	Level              uint32
	Role               uint32
	CryptoPassphase    string
	EnableNotification bool `gorm:"default:true"`
	// UserWallets        []Wallet `gorm:"foreignkey:UserID;auto_preload:true"`
}

// ConvertUserToPbUser : convert user to pb.User
func ConvertUserToPbUser(user *User) *pb.UserInfo {
	return &pb.UserInfo{
		Id:                 user.ID,
		Email:              user.Email,
		Password:           user.Password,
		FullName:           user.FullName,
		Username:           user.UserName,
		Keystore:           user.Keystore,
		EnableNotification: user.EnableNotification,
	}
}

// ConvertPbUserToUser : convert pb.User to User
func ConvertPbUserToUser(user *pb.UserInfo) *User {
	return &User{
		ID:                 user.GetId(),
		Email:              user.GetEmail(),
		Password:           user.GetPassword(),
		FullName:           user.GetFullName(),
		UserName:           user.GetUsername(),
		Keystore:           user.GetKeystore(),
		EnableNotification: user.GetEnableNotification(),
		// UserWallets:        user.GetWallet(),
	}
}
