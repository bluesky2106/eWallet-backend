package mysql

import (
	"testing"

	gwConfig "github.com/bluesky2106/eWallet-backend/config"
	"github.com/bluesky2106/eWallet-backend/entry_store/models"
	"github.com/bluesky2106/eWallet-backend/libs/comparator"
	"github.com/stretchr/testify/assert"
)

var (
	gwConf *gwConfig.Config
	dao    *DAO

	tables = []interface{}{(*models.ProductInfo)(nil), (*models.ProductGroup)(nil), (*models.Unit)(nil)}
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
	unit1 = &models.Unit{
		UID:         "U-01",
		Name:        "Item",
		Description: "",
	}
	unit2 = &models.Unit{
		UID:         "U-02",
		Name:        "Carton",
		Description: "",
	}
)

// product
var (
	prod1 = &models.ProductInfo{
		PID:         "P-01",
		Name:        "T-Shirt",
		Description: "T Shirt for men, size L, color red",
		Size:        models.SizeL,
		Color:       models.Red,
	}
	prod2 = &models.ProductInfo{
		PID:         "P-02",
		Name:        "T-Shirt",
		Description: "T Shirt for men, size M, color red",
		Size:        models.SizeM,
		Color:       models.Red,
	}
)

func init() {
	gwConf = &gwConfig.Config{
		EntryStore: gwConfig.EntryStore{
			MySQL: gwConfig.MySQL{
				Host:     "localhost",
				Port:     "3306",
				Username: "root",
				Password: "Admin123!@#",
				DBName:   "itv_test",
			},
		},
		Env: gwConfig.Debug,
	}
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	var (
		err  error
		conf = ParseConfig(gwConf.EntryStore.MySQL.Username,
			gwConf.EntryStore.MySQL.Password,
			gwConf.EntryStore.MySQL.Host,
			gwConf.EntryStore.MySQL.Port,
			gwConf.EntryStore.MySQL.DBName,
		)
	)

	dao, err = New(conf, string(gwConf.Env))
	assert.Nil(err)
	assert.NotNil(dao)
}

func TestAutoMigrate(t *testing.T) {
	assert := assert.New(t)

	err := dao.AutoMigrate(tables)
	assert.Nil(err)
}

func TestAddForeignKey(t *testing.T) {
	assert := assert.New(t)

	err := dao.AddForeignKey((*models.ProductInfo)(nil), "product_group_id", "product_groups(g_id)")
	assert.Nil(err)
}

func TestCreateProductGroup(t *testing.T) {
	assert := assert.New(t)

	err := dao.WithTransaction(func() error {
		return dao.Create(prodGrp1)
	})
	assert.Nil(err)

	err = dao.WithTransaction(func() error {
		return dao.Create(prodGrp2)
	})
	assert.Nil(err)
}

func TestCreateProductUnit(t *testing.T) {
	assert := assert.New(t)

	err := dao.WithTransaction(func() error {
		return dao.Create(unit1)
	})
	assert.Nil(err)

	err = dao.WithTransaction(func() error {
		return dao.Create(unit2)
	})
	assert.Nil(err)
}

func TestCreateProduct(t *testing.T) {
	assert := assert.New(t)

	// 1. Find ProductGroup whose GID = "PG-01"
	filter := map[string]interface{}{
		"g_id = ?": "PG-01",
	}
	model, err := dao.FindOneByQuery(models.ProductGroup{}, filter)
	assert.Nil(err)
	prodGrp := model.(*models.ProductGroup)

	// 2. Create ProductInfo whose ProductGroupID = the above product group id
	prod1.ProductGroupID = prodGrp.ID
	err = dao.WithTransaction(func() error {
		return dao.Create(prod1)
	})
	assert.Nil(err)

	prod2.ProductGroupID = prodGrp.ID
	err = dao.WithTransaction(func() error {
		return dao.Create(prod2)
	})
	assert.Nil(err)
}

func TestFindOneByQuery(t *testing.T) {
	assert := assert.New(t)

	// 1. Find one product group whose GID = "PG-01"
	filter := map[string]interface{}{
		"g_id = ?": "PG-01",
	}
	// model, err := dao.FindOneByQuery(models.ProductGroup{}, filter, "ProductInfos")
	model, err := dao.FindOneByQuery(&models.ProductGroup{}, filter, "ProductInfos")
	assert.Nil(err)

	// 2. Check ProductInfos inside the above product group
	prodGrp := model.(*models.ProductGroup)
	assert.True(comparator.IsStructValueEqual(prodGrp.ProductInfos[0], prod1, "Model"))
	assert.True(comparator.IsStructValueEqual(prodGrp.ProductInfos[1], prod2, "Model"))
}

func TestFindManyByQuery(t *testing.T) {
	assert := assert.New(t)

	objects, err := dao.FindManyByQuery(models.ProductInfo{}, nil, "ProductGroup")
	// objects, err := dao.FindManyByQuery(&models.ProductInfo{}, nil, "ProductGroup")
	assert.Nil(err)

	products := models.ToProductInfos(objects)
	assert.True(comparator.IsStructValueEqual(products[0], prod1, "Model", "ProductGroup"))
	assert.True(comparator.IsStructValueEqual(products[1], prod2, "Model", "ProductGroup"))

	assert.True(comparator.IsStructValueEqual(products[0].ProductGroup, prodGrp1, "Model", "ProductInfos"))
	assert.True(comparator.IsStructValueEqual(products[1].ProductGroup, prodGrp1, "Model", "ProductInfos"))
}

func TestCountByQuery(t *testing.T) {
	assert := assert.New(t)

	count, err := dao.CountByQuery(models.ProductInfo{}, nil)
	assert.Nil(err)
	assert.Equal(uint(2), count)

	filter := map[string]interface{}{
		"p_id = ?": "P-01",
	}
	count, err = dao.CountByQuery(models.ProductInfo{}, filter)
	assert.Nil(err)
	assert.Equal(uint(1), count)
}

func TestUpdateProduct(t *testing.T) {
	assert := assert.New(t)

	prod1.Description = "Change description"
	err := dao.WithTransaction(func() error {
		return dao.Update(prod1)
	})

	assert.Nil(err)
}

func TestDeleteProduct(t *testing.T) {
	assert := assert.New(t)

	err := dao.WithTransaction(func() error {
		return dao.Delete(prod2)
	})

	assert.Nil(err)
}

func TestDropTables(t *testing.T) {
	assert := assert.New(t)

	err := dao.DropTables(tables...)

	assert.Nil(err)
}
