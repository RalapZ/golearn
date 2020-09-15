package main

import "fmt"

type Stud struct {
	Name string
	Sex  string
	Age  int
}

func (stu Stud) String() string {
	return "Print " + stu.Name + stu.Sex
}

func main() {
	//a := "tesasdfasdfasdfdfasd"
	//b := a[len("tes"):]
	//fmt.Println(b)
	stu := Stud{"ralap", "male", 18}
	//stu.Name= "test"
	fmt.Println(stu)
	fmt.Println(stu.Age)

}
