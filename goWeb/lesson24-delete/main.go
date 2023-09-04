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
	dsn := "zhiwu:123456@tcp(10.79.40.240:33060)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	// 删除符合条件的记录
	var u1 = User{}
	// u1.ID = 5
	// // u1.Name = "jangiu"   //如果不加主键ID，只传Name是无法删除的，会报错 WHERE conditions required
	// db.Debug().Delete(&u1)

	// 批量删除，删除全部匹配的记录
	db.Debug().Where("name=?", "Qing").Delete(&u1)
	// 因为age=40有2条记录，所以不能使用User{}，得使用[]User{}，
	// 否则报错 invalid value, should be pointer to struct or slice
	db.Delete([]User{}, "age=?", 40)

	// 如果使用了 gorm.Model ，在表自动增加了 DeletedAt 字段
	// 这个表在删除记录时没有真正删除记录，而是修改了 DeletedAt 字段
	var u2 []User
	db.Debug().Where("name=?", "xiaoming").Find(&u2)
	fmt.Println(u2)

	// 物理删除，使用了 Unscoped
	db.Debug().Unscoped().Where("name=?", "yun").Delete(User{})
}
