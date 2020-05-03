package model

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
	err error
)

func Init() *gorm.DB {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	dbname := beego.AppConfig.String("mysqldb")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",mysqluser,mysqlpass,url,dbname)
	db,err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败"  + err.Error())
	}
	db.DB().SetMaxOpenConns(20) //数据库连接池中最大连接数
	db.DB().SetMaxIdleConns(10) //连接池中允许最大空闲连接数
	//db.LogMode(true) // 开启sql debug
	return db
}