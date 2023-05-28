package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 测试 http://localhost/web?name=yun&age=1

func main() {
	r := gin.Default()
	r.GET("web", func(ctx *gin.Context) {
		// 获取浏览器请求携带的 query string 参数
		// 方法1：使用 Query
		name := ctx.Query("name") // 通过Query获取请求中携带的querystring参数
		age := ctx.Query("age")
		// 方法2：使用 DefaultQuery
		// name := ctx.DefaultQuery("name", "jiang") // 取不到就用指定默认值
		// 方法3：使用 GetQuery
		// name, ok := ctx.GetQuery("name")
		// if !ok { // 取不到值
		// 	name = "jiang"
		// }
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":80")
}
