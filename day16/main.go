package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status":"ok",
		//})
		// 请求重定向，浏览器地址会改变
		// http重定向，跳转到百度
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {
		// 请求转发，浏览器地址不会改变
		// 跳转到/b对应的路由处理函数
		c.Request.URL.Path = "/b" // 请求修改的URI
		r.HandleContext(c)

	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "bbb",
		})
	})

	r.Run(":9000")
}
