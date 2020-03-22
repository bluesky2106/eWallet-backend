package models

import (
	"github.com/jinzhu/gorm"
)

// User : struct
type User struct {
	gorm.Model
	FullName           string
	UserName           string
	Email              string `gorm:"unique_index"`
	Password           string
	Keystore           string
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

// UserLoginResp : struct
type UserLoginResp struct {
	Token   string `json:"Token"`
	Expired string `json:"Expired"`
}
