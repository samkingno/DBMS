package model

import (
	"singo/util"
	"time"
	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化postgres链接
func Database() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=zhiyinwen dbname=zhiyinwen  sslmode=disable")
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
