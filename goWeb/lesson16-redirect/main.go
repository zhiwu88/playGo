package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://news.sina.cn")
	})

	// /a 内部跳转到 /b
	r.GET("/a", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/b"
		r.HandleContext(ctx)
	})
	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK,hello b",
		})
	})
	r.Run(":80")
}
