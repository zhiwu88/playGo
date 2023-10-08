package routers

import (
	"goweb/lesson26_list_proj/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {})
		// 修改某一个代办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
