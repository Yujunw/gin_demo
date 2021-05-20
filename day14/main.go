package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("name")
		//password := c.Query("pwd")
		//u := User{
		//	username,
		//	password,
		//}
		var u User
		err := c.ShouldBind(&u) // 绑定一个指针
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
		//fmt.Printf("%#v\n", u)
	})

	r.POST("/form", func(c *gin.Context) {
		var u User
		err := c.ShouldBind(&u) // 绑定一个指针
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u User
		err := c.ShouldBind(&u) // 绑定一个指针
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.Run(":9000")

}
