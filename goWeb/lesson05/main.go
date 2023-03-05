package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./templates/hello.html")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	//渲染模板
	user1 := User{
		Name:   "zhiwu88",
		Gender: "man",
		Age:    100,
	}
	map1 := map[string]interface{}{
		"name":   "zhiwu3",
		"gender": "man",
		"age":    99,
	}
	t.Execute(w, map[string]interface{}{
		"u1": user1,
		"m1": map1,
	})
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
