package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	// 连接数据库
	dsn := "zhiwu:123456@tcp(10.79.40.240:33060)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// DB在前面已经声明为全局变量，这里不能在使用 ":="，不然就是重新声明了DB变量
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDB, _ := DB.DB()
	// 要返回的是连接错误信息，并不是连接数数据库 *gorm.DB
	return sqlDB.Ping()
}

func main() {
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	// 模型绑定，创建数据库表
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(ctx *gin.Context) {
			// 从请求中把数据拿出来
			var todo Todo
			ctx.BindJSON(&todo)
			// 存入数据库与返回响应
			// err = DB.Create(&todo).Error
			// if err != nil {
			// 	ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			// }
			if err = DB.Create(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo)
				// ctx.JSON(http.StatusOK, gin.H{
				// 	"code": 2000,
				// 	"msg":  "success",
				// 	"data": todo,
				// })
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(ctx *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {})
		// 修改某一个代办事项
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
				return
			}
			var todo Todo
			if err = DB.Where("id = ?", id).First(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			// 注意这部操作，前端传过来的数据会覆盖todo存着的数据库查询得到的数据
			ctx.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo)
			}

		})
		// 删除某一个代办事项
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
				return
			}
			if err = DB.Where("id = ?", id).Delete(Todo{}).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	r.Run(":80")
}
