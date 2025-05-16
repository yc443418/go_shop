package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TYPE     = "mysql"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	USERNAME = "root"
	PASSWORD = "a625391084."
	PROTOCOL = "tcp"
	DATABASE = "shop"
	CHARSET  = "utf8"
)

var Gdb *gorm.DB

// 中间件是每次路由访问都会执行
func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, PROTOCOL, HOST, PORT, DATABASE, CHARSET)

	var err error
	Gdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // &gorm.Config{} 是 GORM 的配置结构体的一个实例

	if err != nil {
		panic(err)
	}
}
