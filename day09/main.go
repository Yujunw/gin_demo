package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	// 解析模板之前，加载静态文件
	r.Static("/xxx", "./statics")
	//r.LoadHTMLFiles("templates/index.tmpl") // 模板解析
	// gin框架中给模板添加自定义函数，不使用转义
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 通过正则表达式匹配文件，**表示目录，*表示文件
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模板渲染
			"title": "june.com",
		})
	})

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板渲染
			"title": "<a href='https://www.baidu.com'>百度一下，你就知道</a>",
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	// 启动server
	r.Run(":9000")
}
