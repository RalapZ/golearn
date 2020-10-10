package main

import "fmt"

type Stu interface {
	Get()
}

type Stu1 struct {
	Name string
}

type Stu2 struct {
	Test string
}

func (S Stu1) Get() {
	fmt.Println(S.Name)
}

func main() {
	var j Stu
	var stu1 Stu1
	//var stu2 Stu2
	j = stu1
	if k, ok := j.(Stu); ok {
		fmt.Println("this is true")
		fmt.Println(k)
	}
	//v:=stu2
	//if k,ok:=v.(Stu);ok{
	//	fmt.Println("this is true")
	//}

}
