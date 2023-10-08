package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	// 连接数据库
	dsn := "zhiwu:123456@tcp(10.79.40.240:33060)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// DB在前面已经声明为全局变量，这里不能在使用 ":="，不然就是重新声明了DB变量
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDB, _ := DB.DB()
	// 要返回的是连接错误信息，并不是连接数数据库 *gorm.DB
	return sqlDB.Ping()
}
