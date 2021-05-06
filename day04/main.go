package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1. 定义模板hello.tmpl
	// 2. 解析模板
	// 使用相对路径，所以不能直接使用goland的run，找不到./hello.tmpl，要在cli中执行go build
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("template.ParseFiles failed", err)
		return
	}
	// 3. 渲染模板，传入数据data
	err = t.Execute(w, "学习gin")
	if err != nil {
		fmt.Println("render template failed", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe failed", err)
		return
	}
}
