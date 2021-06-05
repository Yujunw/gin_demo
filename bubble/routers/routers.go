package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 在static下寻找静态资源
	r.Static("/static", "static")
	// gin框架在templates下找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1版本
	v1Group := r.Group("v1")
	{
		// 添加待办事项
		v1Group.POST("/todo", controller.CreateATodo)

		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}
