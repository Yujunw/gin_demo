package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET请求url ?后面的是querystring参数
	// key=value格式，多个key-value用&连接
	// eq /web?query=杨超越&age=23
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器那边发送请求的query string
		name := c.Query("query") // 通过Query获取请求中携带的querystring参数
		age := c.Query("age")
		// name := c.DefaultQuery("query", "somebody") // 取不到就用指定的默认值
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9000")
}
