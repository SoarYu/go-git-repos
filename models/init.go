package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 建立数据库连接
func NewDB(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	DB = db
	//err = db.AutoMigrate(&UserBasic{}, &RepoBasic{}, &RepoUser{}, &RepoStar{})
	//if err != nil {
	//	return err
	//}
	return nil
}
