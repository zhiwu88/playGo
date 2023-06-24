package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(ctx *gin.Context) {
		f, err := ctx.FormFile("file1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			dst := path.Join("./", f.Filename)
			ctx.SaveUploadedFile(f, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})
	r.Run(":80")
}
