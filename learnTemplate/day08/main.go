package main

import (
	"log"
	"net/http"
	//"text/template"
	// html/template针对的是需要返回HTML内容的场景，在模板渲染过程中会对一些有风险的内容进行转义，
	// 以此来防范跨站脚本攻击
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板，Go标准库的模板引擎使用的花括号{{和}}作为标识，
	// 而许多前端框架（如Vue和 AngularJS）也使用{{和}}作为标识符，
	// 所以当我们同时使用Go语言模板引擎和以上前端框架时就会出现冲突，
	// 这个时候我们需要修改标识符，修改前端的或者修改Go语言的。
	// 解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		log.Printf("template.ParseFiles failed, %v", err)
		return
	}
	// 渲染模板
	name := "你好"
	err = t.Execute(w, name)
	if err != nil {
		log.Printf("template.Execute failed, %v", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./xss.tmpl")
	if err != nil {
		log.Printf("template.ParseFiles failed, %v", err)
		return
	}
	// 渲染模板，xss攻击
	str := "<script>alert(123);</script>"
	err = t.Execute(w, str)
	if err != nil {
		log.Printf("template.Execute failed, %v", err)
		return
	}
}

func xss2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板之前自定义一个函数safe，表示不需要转义
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		log.Printf("template.ParseFiles failed, %v", err)
		return
	}
	// 渲染模板，xss攻击
	str := "<script>alert(123);</script>"
	err = t.Execute(w, str)
	if err != nil {
		log.Printf("template.Execute failed, %v", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss2)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Printf("http.ListenAndServe failed, %v", err)
		return
	}
}
