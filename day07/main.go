package main

import (
	"log"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		log.Printf("ParseFiles failed, %v", err)
		return
	}

	name := "波波球"
	// ExecuteTemplate方法类似Execute，但是使用名为name的t关联的模板产生输出。
	t.ExecuteTemplate(w, "home.tmpl", name)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		log.Printf("ParseFiles failed, %v", err)
		return
	}

	name := "气球"
	t.ExecuteTemplate(w, "index.tmpl", name)
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalf("http.ListenAndServe failed, %v", err)
	}
}
