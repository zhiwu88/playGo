package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(120);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置字段唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置字段为自增类型
	Address      string  `gorm:"index:addr"`      //给字段创建名为addr的索引
	IgnoreMe     int     `gorm:"_"`               //忽略本字段
}

func main() {
	dsn := "zhiwu:123456@tcp(10.79.40.240:33060)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 创建数据库表
	db.AutoMigrate(&User{})
}
