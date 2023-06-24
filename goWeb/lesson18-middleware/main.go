package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(ctx *gin.Context) {
	fmt.Println("index")
	name, ok := ctx.Get("name") // 从上下文中获取值
	if !ok {
		name = "匿名用户"
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": name})
}

// 中间件可用于登陆认证、VIP会员、记录日志、耗时统计
func mid1(ctx *gin.Context) {
	fmt.Println("middleware1 in ...")
	// 计时，统计耗时
	start := time.Now()
	ctx.Next() // 调用后续的处理函数
	// ctx.Abort()  // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("middleware1 out ...")
}

func mid2(ctx *gin.Context) {
	fmt.Println("middleware2 in ...")
	ctx.Set("name", "Jim") // 在上下文中设置值
	ctx.Next()             // 调用后续的处理函数
	fmt.Println("middleware2 out ...")
}

func mid3(ctx *gin.Context) {
	fmt.Println("middleware3 in ...")
	ctx.Next() // 调用后续的处理函数
	fmt.Println("middleware3 out ...")
}

func main() {
	r := gin.Default()
	// r.GET("/index", mid1, indexHandler)
	// r.GET("/user", mid1, func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "user"}) })

	// 使用全局注册中间件后，不再需要在每个路由处理上添加中间件函数名。
	//r.Use(mid1)       // 全局注册中间件函数mid1
	r.Use(mid1, mid2) // 全局注册中间件函数mid1、mid2
	r.GET("/index", indexHandler)
	r.GET("/user", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "user"}) })

	// 路由组使用中间件有2种方法，另外上面的全局注册同样作用于路由组
	// videoGroup := r.Group("/video", mid3)
	// {
	// 	videoGroup.GET("/index", indexHandler)
	// }
	videoGroup := r.Group("/video")
	videoGroup.Use(mid3)
	{
		videoGroup.GET("/index", indexHandler)
	}

	r.Run(":80")
}
