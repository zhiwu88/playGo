package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(ctx *gin.Context) {
		// 使用 PostForm
		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password")
		// 使用 DefaultPostForm，取不到返回默认值
		// username := ctx.DefaultPostForm("username", "zhangsan")
		// password := ctx.DefaultPostForm("password", "123")
		// 使用 GetPostForm，根据是否取到返回 true/false
		username, ok := ctx.GetPostForm("username")
		if !ok {
			username = "lisi"
		}
		password, _ := ctx.GetPostForm("password")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":80")
}
