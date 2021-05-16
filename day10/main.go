package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// gin框架返回json，前后端分离，不需要返回模板，后端只需要返回相应的json
	r.GET("/json", func(c *gin.Context) {
		// 1. 直接传入map
		//data := map[string]interface{}{
		//	"name":    "June",
		//	"message":    "coder",
		//	"age":    "26",
		//}
		// 2. 使用gin.H
		data := gin.H{"name": "June", "message": "coder", "age": "26"}
		c.JSON(http.StatusOK, data)
	})

	// 3. 使用结构体，灵活使用tag
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"june",
			"learn gin",
			26,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9000")
}
