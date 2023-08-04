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

	var user1 User
	db.First(&user1)
	user1.Name = "Jangiu"
	user1.Age = 27
	db.Save(&user1)                        // 默认修改所有字段
	db.Model(&user1).Update("name", "yun") // 使用Update只更新单项数据

	var user2 User
	db.Last(&user2)
	map1 := map[string]interface{}{
		"name": "bangyu",
		"age":  88,
	}
	// db.Model(&user2).Updates(map1) // 使用 map 更新多个属性，只会更新其中有变化的属性
	// db.Model(&user2).Select("age").Updates(map1) // 使用 Select 选择只更新某个字段
	db.Model(&user2).Omit("age").Updates(map1) // 使用 Omit 排除不需要更新的字段

	// 给表字段age的所有值批量增加2
	var user3 []User
	db.Debug().Find(&user3)
	fmt.Printf("data: %v/n", user3)
	db.Model(&user3).Update("age", gorm.Expr("age+?", 2))
}
