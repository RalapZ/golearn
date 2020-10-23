package main

import (
	"os"
	"text/template"
)

func SayHello() string {
	return "hello world"
}

func SayYouName(name string) string {
	return "hello" + name
}

func demo2() {
	k := func(name string) (string, error) {
		return name + "年轻帅气", nil
	}

	t := template.New("t.tmpl")
	t.Funcs(template.FuncMap{
		"test2": k,
		//"test" : "
	})
	t.ParseFiles("src/base/gintest/template/t.tmpl")
	name := "my zone"
	t.Execute(os.Stdout, name)

}
func main() {
	funcMap := template.FuncMap{
		//在FuncMap中声明相应要使用的函数，然后就能够在template字符串中使用该函数
		"SayHello":   SayHello,
		"SayYouName": SayYouName,
	}
	//name := "boss"
	//tmpl, err := template.New("test").Funcs(funcMap).Parse("{{SayHello}}\n{{SayYouName .}}\n")

	t := template.New("test")
	t.Funcs(funcMap)
	demo2()
	t.Parse("{{ SayHello }}  {{ SayYouName . }}")
	//template.New("test").Funcs(funcMap).Parse("{{SayHello}}\n{{SayYouName .}}\n")
	//err = tmpl.Execute(os.Stdout, name)
	//err := t.Execute(os.Stdout,name)
	//if err != nil{
	//	panic(err)
	//}

}
