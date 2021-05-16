package database

import (
	"bookShop.seven.com/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Conn *gorm.DB

func init() {

	//根据 config配置来
	db, err := gorm.Open("mysql", conf.MySqlRoot)

	if err != nil {
		fmt.Print("初始化数据连接失败，失败原因：")
		panic(err)
	}
	Conn = db
	//初始化连接池
	initPool(20, 130)
	// 启用Logger，显示详细日志
	setLogMode(true)
}

func initPool(maxIdleConns, maxOpenConns int) {
	//SetMaxIdleConns用于设置闲置的连接数
	//SetMaxOpenConns用于设置最大打开的连接数
	Conn.DB().SetMaxIdleConns(maxIdleConns)
	Conn.DB().SetMaxOpenConns(maxOpenConns)
}

func setLogMode(status bool) {
	// 启用Logger，显示详细日志
	Conn.LogMode(status)
}
