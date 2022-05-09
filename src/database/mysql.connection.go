package database

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-role/src/config"
	"github.com/samithiwat/samithiwat-backend-role/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(conf *config.Database) (gormDb *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", conf.User, conf.Password, conf.Host, conf.Name)

	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = gormDb.AutoMigrate(model.Permission{}, model.Role{})
	if err != nil {
		return
	}

	return
}
