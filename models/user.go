package models

import (
	"github.com/jinzhu/gorm"
)

// User : struct
type User struct {
	ID                 uint
	FullName           string
	UserName           string
	Email              string `gorm:"unique_index"`
	Password           string
	Keystore           string
	CryptoPassphase    string
	EnableNotification bool `gorm:"default:true"`
	// UserWallets        []Wallet `gorm:"foreignkey:UserID;auto_preload:true"`
}

// UserPassword : struct
type UserPassword struct {
	gorm.Model
	User   *User
	UserID uint

	Code    string
	Expired uint64
	Retry   uint `gorm:"default:0"`
}

// UserRegisterReq : user register request
type UserRegisterReq struct {
	FullName        string `json:"FullName"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}

// UserLoginReq : struct
type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLoginResp : struct
type UserLoginResp struct {
	Token   string `json:"Token"`
	Expired string `json:"Expired"`
}
