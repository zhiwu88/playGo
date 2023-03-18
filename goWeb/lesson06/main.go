package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//测试模板嵌套

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数
	kua := func(name string) (string, error) {
		return name + "帅！", nil
	}
	//定义模板
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{"kua99": kua})

	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	name := "zhiwu88"
	//渲染模板
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	name := "zhiwu88888"
	//渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpldemo", demo1)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Http server start failed, err: %v", err)
		return
	}
}
