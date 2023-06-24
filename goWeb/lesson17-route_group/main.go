package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// HTTP请求的方法有很多 GET/POST/PUT/DELETE/HEAD ... 逐个写如下。
	// 测试方法 curl -X PUT http://localhost/method
	r.GET("/method", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"Method": "GET"}) })
	r.POST("/method", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"Method": "POST"}) })
	r.PUT("/method", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"Method": "PUT"}) })
	r.DELETE("/method", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"Method": "DELETE"}) })

	// 使用Any可以处理所有方法的请求
	r.Any("/any", func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case "GET":
			ctx.JSON(http.StatusOK, gin.H{"Method": "GET"})
		case http.MethodPost:
			ctx.JSON(http.StatusOK, gin.H{"Method": "POST"})
		case http.MethodPut:
			ctx.JSON(http.StatusOK, gin.H{"Method": "PUT"})
		}
	})

	// 使用NoRoute统一接处理不存在的路由（URL路径）
	r.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://sina.cn")
	})

	// 使用路由组组织管理URL，把公用的前缀提取出来创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "url is /video/index"}) })
		videoGroup.GET("/xx", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "url is /video/xx"}) })
		videoGroup.GET("/oo", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "url is /video/oo"}) })
	}

	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "url is /shop/index"}) })
		// 嵌套路由组
		xx := shopGroup.Group("/xx")
		{
			xx.GET("/oo", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "url is /shop/xx/oo"}) })
		}
	}

	r.Run(":80")
}
