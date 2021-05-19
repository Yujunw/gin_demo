package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 注意URL的匹配不要冲突
	r.GET("/:name/:age", func(c *gin.Context) {
		// 获取路径参数，请求返回的都是string类型
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})

	r.Run(":9000")
}
