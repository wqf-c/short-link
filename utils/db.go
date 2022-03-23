package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"short-link/common"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	//dsn := "root:123456@tcp(localhost:3306)/short-link"
	dsn := common.Info.Mysql.Url
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}
