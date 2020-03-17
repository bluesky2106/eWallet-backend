package mysql

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// DAO struct
type DAO struct {
	db   *gorm.DB
	conf *Config
}

// New : connect mysql server
func New(conf *config.Config) (*DAO, error) {
	sqlConf := &Config{
		DBName:   conf.MySQL.DBName,
		Host:     conf.MySQL.Host,
		Port:     conf.MySQL.Port,
		Username: conf.MySQL.Username,
		Password: conf.MySQL.Password,
	}
	env := conf.Env

	connURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		sqlConf.Username, sqlConf.Password, sqlConf.Host, sqlConf.Port, sqlConf.DBName)
	db, err := gorm.Open("mysql", connURL)
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}
	if env == config.Production {
		db.LogMode(true)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci auto_increment=1")
	// skip save associations of gorm -> manual save by code
	db = db.Set("gorm:save_associations", false)
	db = db.Set("gorm:association_save_reference", true)
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(10)

	return &DAO{
		conf: sqlConf,
		db:   db,
	}, nil
}
