package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 使用map
	r.GET("/json", func(ctx *gin.Context) {
		// 使用自定义map
		// data := map[string]interface{}{
		// 	"name":    "云卿",
		// 	"message": "I love you, baby!",
		// }
		// 使用内置 gin.H 类型
		data := gin.H{
			"name":    "云卿",
			"message": "I love you, baby!",
		}
		ctx.JSON(http.StatusOK, data)
	})

	// 使用结构体
	type msg struct {
		Name    string
		Message string `json:"msg"` //可以灵活使用tag对结构体字段定制化操作。该tag意思时在json格式化时，字段名称改成msg
	}
	r.GET("/jsons", func(ctx *gin.Context) {
		data := msg{
			Name:    "yunqing",
			Message: "I love you",
		}
		ctx.JSON(http.StatusOK, data)
	})

	r.Run(":80")
}
