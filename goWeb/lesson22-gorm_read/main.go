package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
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

	// 创建数据
	// u1 := User{Name: "xiaoming", Age: 18}
	// db.Create(&u1)
	// u2 := User{Name: "dazhuang", Age: 18}
	// db.Create(&u2)

	// 查询数据
	var user User
	db.First(&user) // 查询第一条数据
	fmt.Printf("first row:%#v\n", user)

	var user2 User
	db.Last(&user2) // 查询最后一条数据
	fmt.Printf("last row: %v\n", user2)

	// 查询结果 与 操作返回值 不一样
	var users []User
	result := db.Debug().Find(&users) // 查询所有数据，数据存在 users，操作返回信息存在 result
	fmt.Printf("all rows: %#v\n", users)
	fmt.Printf("rows count:%v\n", result.RowsAffected)

	// where 条件查询
	var users2 []User
	db.Debug().Where("name like ?", "%Qing%").Find(&users2)
	fmt.Printf("all rows: %v\n", users2)

	// Or 查询条件
	var users3 []User
	db.Debug().Where("name like ?", "%Qing%").Or("age = ?", 28).Find(&users3)
	fmt.Printf("all rows: %v\n", users3)

	// 使用 Struct 查询
	var users4 []User
	db.Debug().Where("name like ?", "%Qing%").Or(User{Age: 28}).Find(&users4)
	fmt.Printf("all rows: %v\n", users4)

	// 内联条件 ， 与 where 类似
	var users5 []User
	db.Debug().Find(&users5, "name like ? OR age = ?", "%Qing%", 28)
	fmt.Printf("all rows: %v\n", users5)

	// 使用给定条件查询，查不到就把给定条件赋给变量
	var user3 User
	db.Debug().FirstOrInit(&user3, User{Name: "Jame", Age: 18})
	fmt.Printf("all rows: %v\n", user3)

	// 使用 Select 拿指定字段
	var users6 []User
	db.Debug().Select("name,age").Where("name like ?", "%Qing%").Or("age = ?", 28).Find(&users6)
	fmt.Printf("all rows: %v\n", users6)

	// 使用 Order 排序
	var users7 []User
	db.Debug().Order("age desc").Where("name like ?", "%Qing%").Or("age = ?", 28).Find(&users7)
	fmt.Printf("all rows: %v\n", users7)

	// 使用 Limit 限制数量
	var users8 []User
	db.Debug().Limit(2).Where("name like ?", "%Qing%").Or("age = ?", 28).Find(&users8)
	fmt.Printf("all rows: %v\n", users8)

	// https://www.liwenzhou.com/posts/Go/gorm-crud/
}
