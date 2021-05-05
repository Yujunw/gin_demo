package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"hello golang",
	})
}

func main() {
	// 返回默认的路由引擎
	r := gin.Default()
	// 指定用户使用GET请求访问/hello时，执行sayHello函数
	r.GET("/hello", sayHello)

	// 启动服务，默认在8080端口
	r.Run(":9090")
}
