package mysql

import (
	"testing"

	"github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/models"
	"github.com/stretchr/testify/assert"
)

var (
	conf *config.Config
	dao  *DAO

	tables = []interface{}{(*models.Product)(nil), (*models.ProductGroup)(nil), (*models.ProductUnit)(nil)}
)

// product group
var (
	prodGrp1 = &models.ProductGroup{
		GID:         "PG-01",
		Name:        "Ladies fashion",
		Description: "Clothes for women",
	}
	prodGrp2 = &models.ProductGroup{
		GID:         "PG-02",
		Name:        "Men 's fashion",
		Description: "Clothes for men",
	}
)

// product unit
var (
	prodUnit1 = &models.ProductUnit{
		UID:         "PU-01",
		Name:        "Item",
		Description: "",
	}
	prodUnit2 = &models.ProductUnit{
		UID:         "PU-02",
		Name:        "Carton",
		Description: "",
	}
)

// product
var (
	prod1 = &models.Product{
		PID:            "P-01",
		Name:           "T-Shirt",
		Description:    "T Shirt for men, size L, color red",
		Size:           models.SizeL,
		Color:          models.Red,
		ProductGroupID: "PG-01",
		ProductUnitID:  "PU-01",
		CostPerUnit:    100,
	}
	prod2 = &models.Product{
		PID:            "P-02",
		Name:           "T-Shirt",
		Description:    "T Shirt for men, size M, color red",
		Size:           models.SizeM,
		Color:          models.Red,
		ProductGroupID: "PG-01",
		ProductUnitID:  "PU-01",
		CostPerUnit:    110,
	}
)

func init() {
	conf = &config.Config{
		MySQL: config.MySQL{
			Host:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "itv",
			DBName:   "itv_test",
		},
		Env: config.Debug,
	}
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	var err error
	dao, err = New(conf)
	assert.Nil(err)
	assert.NotNil(dao)
}

func TestAutoMigrate(t *testing.T) {
	assert := assert.New(t)

	err := dao.AutoMigrate(tables)
	assert.Nil(err)
}

func TestAddForeignKey(t *testing.T) {
	// db.Model(&models.Wallet{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	assert := assert.New(t)

	err := dao.AddForeignKey((*models.Product)(nil), "product_group_id", "product_groups(g_id)")
	assert.Nil(err)

	err = dao.AddForeignKey((*models.Product)(nil), "product_unit_id", "product_units(u_id)")
	assert.Nil(err)
}

func TestCreateProductGroup(t *testing.T) {
	assert := assert.New(t)

	err := dao.WithTransaction(func() error {
		return dao.Create(prodGrp1)
	})
	err = dao.WithTransaction(func() error {
		return dao.Create(prodGrp2)
	})

	assert.Nil(err)
}

func TestCreateProductUnit(t *testing.T) {
	assert := assert.New(t)

	err := dao.WithTransaction(func() error {
		return dao.Create(prodUnit1)
	})

	err = dao.WithTransaction(func() error {
		return dao.Create(prodUnit2)
	})

	assert.Nil(err)
}

func TestCreateProduct(t *testing.T) {
	assert := assert.New(t)

	err := dao.WithTransaction(func() error {
		return dao.Create(prod1)
	})

	err = dao.WithTransaction(func() error {
		return dao.Create(prod2)
	})

	assert.Nil(err)
}

func TestDropTables(t *testing.T) {
	assert := assert.New(t)

	err := dao.DropTables(tables...)

	assert.Nil(err)
}
