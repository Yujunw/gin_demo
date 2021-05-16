package main

import (
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	Gender string // 小写，无法导出
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Printf("template.ParseFiles failed, err = %v", err)
		return
	}
	// 渲染模板
	u1 := User{
		Name:   "喻君武",
		Gender: "男",
		Age:    26,
	}
	m1 := map[string]interface{}{
		"name":   "喻君武",
		"age":    26,
		"gender": "男",
	}

	hobbyList := []string{
		"抽烟",
		"喝酒",
		"烫头",
	}
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Printf("http.ListenAndServe failed, err = %v", err)
	}
}
