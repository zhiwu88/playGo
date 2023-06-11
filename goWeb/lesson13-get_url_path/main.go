package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(ctx *gin.Context) { //把路径赋值给指定名
		// 获取路径参数
		name := ctx.Param("name")
		age := ctx.Param("age") // string类型
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":80")
}
