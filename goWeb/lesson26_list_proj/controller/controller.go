package controller

import (
	"goweb/lesson26_list_proj/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
url --> controller --> logic --> model
请求 --> 控制器 --> 业务逻辑 --> 模型层的增删改查
*/

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(ctx *gin.Context) {
	// 从请求中把数据拿出来
	var todo models.Todo
	ctx.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(ctx *gin.Context) {
	todoList, err := models.GetAllTodo()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 注意这部操作，前端传过来的数据会覆盖todo存着的数据库查询得到的数据
	ctx.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}

}

func DeleteATodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
