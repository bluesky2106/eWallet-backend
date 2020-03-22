package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

// ProductGroup : product group struct
type ProductGroup struct {
	gorm.Model
	GID          string `gorm:"unique_index"`
	Name         string `gorm:"unique_index"`
	Description  string
	ProductInfos []*ProductInfo
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

// ProductInfo : product information
type ProductInfo struct {
	gorm.Model
	PID            string `gorm:"unique_index"`
	Name           string
	Description    string
	ProductGroupID uint
	ProductGroup   *ProductGroup
	Size           ProductSize  `gorm:"column:size" json:"size"`
	Color          ProductColor `gorm:"column:color" json:"color"`
}

// ProductInfos : slice of *ProductInfo
type ProductInfos []*ProductInfo

// ToProductInfos converts interface to slice of products
func ToProductInfos(objects interface{}) ProductInfos {
	list := reflect.ValueOf(objects)
	products := make(ProductInfos, list.Len())
	for i := 0; i < list.Len(); i++ {
		products[i] = list.Index(i).Interface().(*ProductInfo)
	}

	return products
}
