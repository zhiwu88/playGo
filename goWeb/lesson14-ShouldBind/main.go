package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`  //使用tag对结构体字段定制化操作。
	Password string `form:"password" json:"pwd"`   //该tag意思时在from格式化时，字段名称改成password；json格式化时，字段名称改成pwd
}

func main() {
	r := gin.Default()

	// http://localhost/user?username=zhiwu&password=123
	r.GET("/user", func(ctx *gin.Context) {
		// username := ctx.Query("username")
		// password := ctx.Query("password")
		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	// curl --data "username=zhiwu&password=123456" http://localhost/form
	r.POST("/form", func(ctx *gin.Context)  {
		// username := ctx.PostForm("username")
		// password := ctx.PostForm("password")
		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	// curl -H "Content-Type: application/json" --data '{"user":"zhiwu","pwd":"123456"}' http://localhost/json
	r.POST("/json", func(ctx *gin.Context)  {
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.Run(":80")
}
