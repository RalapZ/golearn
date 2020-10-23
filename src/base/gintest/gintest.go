package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Name string
	Age  int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("src/base/gintest/hello.tmpl")
	stu := Student{"ralap.L", 22}
	std2 := Student{"ralap.Z", 18}
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	test1 := map[string]Student{
		"key1": stu,
		"key2": std2,
	}

	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, test1)
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
