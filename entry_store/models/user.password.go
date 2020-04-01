package models

import (
	"github.com/jinzhu/gorm"
)

// UserPassword : user - password relationship is 1-n
type UserPassword struct {
	gorm.Model
	User   *User
	UserID uint

	Code    string
	Expired uint64
	Retry   uint `gorm:"default:0"`
}
