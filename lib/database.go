package lib

import (
	"fmt"
	"invitnesia/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDBConnection(cfg *config.Database) (err error) {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               DSN,
		DefaultStringSize: 256,
	}), &gorm.Config{})

	return err
}
