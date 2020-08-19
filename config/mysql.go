package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", MysqlUrl)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.LogMode(true) //打印SQL日志
	if err != nil {
		panic("数据库连接失败")
	}
}
