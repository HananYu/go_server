package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:921115@tcp(111.230.12.75:3306)/yusj?charset=utf8&parseTime=true")
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.LogMode(true) //打印SQL日志
	if err != nil {
		panic("数据库连接失败")
	}
}
