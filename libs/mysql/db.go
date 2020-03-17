package mysql

import (
	"fmt"

	"github.com/bluesky2106/eWallet-backend/config"
	errs "github.com/bluesky2106/eWallet-backend/errors"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
		return nil, errs.New(errs.ECMySQLConnection, err.Error(), "gorm.Open")
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

// DB returns dao.db
func (dao *DAO) DB() *gorm.DB {
	return db
}

// WithTransaction : commit transaction
func (dao *DAO) WithTransaction(callback func() error) error {
	tx := dao.db.Begin()

	if err := callback(); err != nil {
		tx.Rollback()
		return errs.WithMessage(err, "callback")
	}

	if err := tx.Commit().Error; err != nil {
		return errs.WithMessage(err, "tx.Commit()")
	}

	return nil
}

// AutoMigrate : initiate tables
func (dao *DAO) AutoMigrate(tables []interface{}) error {
	if dao.db == nil {
		return emptyDBError("db.AutoMigrate")
	}

	if err := dao.db.AutoMigrate(tables...).Error; err != nil {
		return errs.New(errs.ECMySQLDBAutoMigrate, err.Error(), "db.AutoMigrate")
	}

	return nil
}

// AddForeignKeys : foreign keys
func (dao *DAO) AddForeignKeys(model interface{}, field, dest string) error {
	if dao.db == nil {
		return emptyDBError("db.AddForeignKeys")
	}

	dao.db.Model(model).AddForeignKey(field, dest, "CASCADE", "CASCADE")

	return nil
}

// Create : create model
func (dao *DAO) Create(model interface{}) error {
	err := dao.db.Create(model).Error
	if err != nil {
		return errs.New(errs.ECMySQLCreate, err.Error(), "db.Create")
	}
	return nil
}

// Update : update model
func (dao *DAO) Update(model interface{}) error {
	err := dao.db.Save(model).Error
	if err != nil {
		return errs.New(errs.ECMySQLUpdate, err.Error(), "db.Save")
	}
	return nil
}

// FindAll : find all objects matched the filters
func (dao *DAO) FindAll(filters map[string]interface{}) (interface{}, error) {
	var (
		objects interface{}
		db      = dao.db
	)
	for k, v := range filters {
		if v != nil {
			db = db.Where(k, v)
		} else {
			db = db.Where(k)
		}
	}

	if err := db.Find(&objects).Error; err != nil {
		return nil, errs.New(errs.ECMySQLRead, err.Error(), "db.FindAll")
	}
	return objects, nil
}

// DeleteAll : delete all models matched the filters
func (dao *DAO) DeleteAll(model interface{}, filters map[string]interface{}) error {
	db := dao.db

	for k, v := range filters {
		if v != nil {
			db = db.Where(k, v)
		} else {
			db = db.Where(k)
		}
	}

	err := db.Delete(model).Error
	if err != nil {
		return errs.New(errs.ECMySQLDelete, err.Error(), "db.Delete")
	}

	return nil
}

// CountByQuery : count number of objects matched the filters
func (dao *DAO) CountByQuery(filters map[string]interface{}) (uint, error) {
	var count uint

	query := dao.db
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v)
		} else {
			query = query.Where(k)
		}
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, errs.New(errs.ECMySQLRead, err.Error(), "query.Count")
	}
	return count, nil
}

// DropTable : drop table if it is exist
func (dao *DAO) DropTable(models ...interface{}) error {
	if err := dao.db.DropTableIfExists(models...).Error; err != nil {
		return errs.New(errs.ECMySQLDelete, err.Error(), "db.DropTableIfExists")
	}
	return nil
}
