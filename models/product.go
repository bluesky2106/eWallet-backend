package models

import "github.com/jinzhu/gorm"

// ProductGroup : product group struct
type ProductGroup struct {
	gorm.Model
	GID         string `gorm:"unique_index"`
	Name        string `gorm:"unique_index"`
	Description string
	Products    []*Product
}

// ProductUnit : calculation unit
type ProductUnit struct {
	gorm.Model
	UID         string `gorm:"unique_index"`
	Name        string `gorm:"unique_index"`
	Description string
}

// ProductSize : this type is useful for clothes (ex. size S, M, L, XL, etc)
type ProductSize string

const (
	// SizeL : size L
	SizeL ProductSize = "Size L"
	// SizeM : size M
	SizeM ProductSize = "Size M"
	// SizeXL : size XL
	SizeXL ProductSize = "Size XL"
)

// ProductColor : enum for Product color
type ProductColor string

const (
	// Red :
	Red ProductColor = "Red"
)

// Product : product information
type Product struct {
	gorm.Model
	PID            string `gorm:"unique_index"`
	Name           string `gorm:"column:name" json:"name"`
	Description    string `gorm:"column:description" json:"description"`
	ProductGroupID string `gorm:"column:product_group_id" json:"product_group_id"`
	ProductGroup   *ProductGroup
	Size           ProductSize  `gorm:"column:size" json:"size"`
	Color          ProductColor `gorm:"column:color" json:"color"`
	ProductUnitID  string       `gorm:"column:product_unit_id" json:"product_unit_id"`
	ProductUnit    *ProductUnit
	CostPerUnit    float64 `gorm:"column:cost_per_unit" json:"cost_per_unit"`
}
