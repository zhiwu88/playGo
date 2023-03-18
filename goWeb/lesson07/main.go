package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//测试模板继承

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	//渲染模板
	msg := "这是index页面"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("templates/home.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	//渲染模板
	msg := "这是home页面"
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("templates/base.tmpl", "templates/index2.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	//渲染模板
	msg := "这是index2页面"
	t.Execute(w, msg)
}
func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("templates/base.tmpl", "templates/home2.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	//渲染模板
	msg := "这是home2页面"
	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Http server start failed, err: %v", err)
		return
	}
}
