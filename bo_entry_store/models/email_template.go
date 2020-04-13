package models

import (
	"github.com/jinzhu/gorm"
)

// EmailTemplate : struct
type EmailTemplate struct {
	gorm.Model
	SendgridTemplateID string `json:"sendgrid_template_id"`
	Type               uint   `json:"type" gorm:"unique;not null"`
}
