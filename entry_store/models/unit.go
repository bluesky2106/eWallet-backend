package models

import "github.com/jinzhu/gorm"

// Unit struct
type Unit struct {
	gorm.Model
	UID         string `gorm:"unique_index"`
	Name        string `gorm:"unique_index"`
	Description string
}
