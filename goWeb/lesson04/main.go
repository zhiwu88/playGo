package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./templates/hello.html")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	//渲染模板
	name := "zhiwu88"
	t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err: %v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v", err)
		return
	}
}
