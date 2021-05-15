package main

import (
	"log"
	"net/http"
	"text/template"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 定义一个函数
	k := func(name string) (string, error) {
		return name + "真好看！", nil
	}

	// 解析模板
	// 创建一个模板对象，名称与模板文件名称要相同
	t := template.New("hello.tmpl")
	// 告诉模板引擎，多了一个自定义函数
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	_, err := t.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Printf("ParseFiles failed, %v", err)
	}
	// 渲染模板
	name := "波波球"
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 父模板在前，子模板在后
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		log.Printf("ParseFiles failed, %v", err)
	}

	name := "波波球"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpl", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Printf("http.ListenAndServe failed, %v", err)
	}
}
