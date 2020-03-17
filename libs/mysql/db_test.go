package mysql

import (
	"testing"

	"github.com/bluesky2106/eWallet-backend/config"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	conf *config.Config
	dao  *DAO
)

func init() {
	conf = &config.Config{
		MySQL: config.MySQL{
			Host:     "localhost",
			Port:     "3307",
			Username: "tokoin",
			Password: "tokoin",
			DBName:   "tokoin",
		},
		Env: config.Production,
	}
}

type Person struct {
	gorm.Model
	Name string `gorm:"unique_index"`
	Age  uint8
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

	tables := []interface{}{(*Person)(nil)}
	err := dao.AutoMigrate(tables)
	assert.Nil(err)
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	akagi := &Person{
		Name: "Akagi",
		Age:  uint8(34),
	}

	err := dao.WithTransaction(func() error {
		return dao.Create(akagi)
	})

	assert.Nil(err)
}

func TestDropTable(t *testing.T) {
	assert := assert.New(t)

	akagi := &Person{
		Name: "Akagi",
		Age:  uint8(34),
	}

	err := dao.WithTransaction(func() error {
		return dao.DropTable(akagi)
	})

	assert.Nil(err)
}
