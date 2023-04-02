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
	r.Static("/xxx", "./statics")
	//模板解析
	//r.LoadHTMLFiles("templates/index.html")
	r.LoadHTMLGlob("templates/**/*")

	// r.GET("/index", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.html", gin.H{ //模板渲染
	// 		"title": "Hello gin template.",
	// 	})
	// })

	r.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.html", gin.H{ //模板渲染
			"title": "Hello posts/index.",
		})
	})

	r.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.html", gin.H{ //模板渲染
			"title": "Hello users/index.",
		})
	})

	r.GET("/users/a", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.html", gin.H{ //模板渲染
			"title": "<a href='https://sina.cn'>大浪</a>",
		})
	})

	r.Run(":80")
}
