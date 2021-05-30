package main

// 中间件，Gin中的中间件必须是一个gin.HandlerFunc类型。

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 也是HandlerFunc类型
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"path": "index",
	})
}

// 定义一个中间件类型
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 计时
	start := time.Now()
	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost%v\n", cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Next()
	fmt.Println("m2 out ...")
}

func main() {
	// gin.Default()默认使用了Logger和Recovery中间件，其中：
	//Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	//Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	//如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

	//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），
	//必须使用其只读副本（c.Copy()
	r := gin.Default()
	// 全局注册中间件m1, m2
	r.Use(m1, m2)
	r.GET("/index", indexHandler)

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"path": "/shop"})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"path": "/user"})
	})

	r.Run(":9000")
}
