package main

import (
	"html/template"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {

	k := func(name string) (string, error) {
		return name + "年轻帅气", nil
	}

	t := template.New("t.tmpl")
	t.Funcs(template.FuncMap{
		"test2": k,
	})
	t.ParseFiles("src/base/gintest/template/t.tmpl")
	name := "my zone"

	t.Execute(w, name)

}
func templ(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("src/base/gintest/template/tmp/t2.tmpl", "src/base/gintest/template/tmp/ur.tmpl")

	name := "my zone"

	tmpl.Execute(w, name)

}
func main() {
	http.HandleFunc("/", demo)
	http.HandleFunc("/tmpl", templ)
	http.ListenAndServe(":9091", nil)
}
