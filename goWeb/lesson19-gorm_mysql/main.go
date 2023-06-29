package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// mysql -h 10.79.40.240 -u zhiwu -P 33060 -p
	// create database db1;  手工创建数据库
	// grant all privileges on *.* to 'zhiwu'@'%';  授权访问所有数据库
	// db, err := gorm.Open("mysql", "zhiwu:123456@(10.79.40.240:33060)/db1?charset=utf8mb48&parseTime=True&loc=Local") //老方法
	dsn := "zhiwu:123456@tcp(10.79.40.240:33060)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// defer db.Close
	// 创建表，自动迁移，把结构体和数据表进行对应
	db.AutoMigrate(&UserInfo{})
	// 创建数据行
	u1 := UserInfo{1, "Yun", "girl", "Music"}
	db.Create(&u1)
	// 查询数据行
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)
	// 更新数据行
	db.Model(&u).Update("hobby", "study")
	db.First(&u)
	fmt.Printf("u:%#v\n", u)
	// 删除数据行
	db.Delete(&u)
}
