package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello Gin!",
	})
}

func main() {
	r := gin.Default() //返回默认的路由引擎

	//使用无返回值普通函数，通过修改指针参数的值来传值到函数外部
	r.GET("/hello", sayHello)

	//使用匿名函数，也是通过修改指针参数的值来传值到函数外部
	r.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.Run(":80")
}
