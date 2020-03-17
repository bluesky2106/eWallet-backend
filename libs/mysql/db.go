package mysql

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Init : connect mysql server
func Init(conf *config.MySQL, env config.Environment) (*gorm.DB, error) {
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
	db, err := gorm.Open("mysql", connURL)
	if env == config.Production {
		db.LogMode(true)
	}
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci auto_increment=1")
	// skip save associations of gorm -> manual save by code
	db = db.Set("gorm:save_associations", false)
	db = db.Set("gorm:association_save_reference", true)
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(10)
	return db, err
}
