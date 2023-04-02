package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//gin框架中给模板添加自已定义函数，用于防止渲染内容中的HTML标记被转义成文本
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//加载静态文件
	r.Static("/static", "./statics")
	//模板解析
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/wallet.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "wallet.html", nil)
	})

	r.Run(":80")
}
