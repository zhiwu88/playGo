package main

import (
	"goweb/lesson26_list_proj/dao"
	"goweb/lesson26_list_proj/models"
	"goweb/lesson26_list_proj/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 模型绑定，创建数据库表
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":80")
}
