package main

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int64
	Name *string `gorm:"default:'Qing'"` // 在string前加*，变成指针类型，解决无法写入空值的问题
	Age  int64
	Addr sql.NullString `gorm:"default:'ZPARK'"` // 使用 Null+String 解决无法写入空值的问题
}

func main() {
	// 连接数据库
	dsn := "zhiwu:123456@tcp(10.79.40.240:33060)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	// 创建记录
	// u := User{Name: "bang", Age: 28}
	u1 := User{Age: 18}
	db.Create(&u1)

	// *string 类型，使用 new(string) 设置空值，因为 new(）函数返回的是指针类型
	u2 := User{Name: new(string), Age: 28}
	db.Create(&u2)

	// sql.NullString 类型，使用 sql.NullString{String: "", Valid: true} 设置空值
	strName := "XiaoMing"
	u3 := User{Name: &strName, Age: 38, Addr: sql.NullString{String: "", Valid: true}}
	db.Create(&u3)
}
