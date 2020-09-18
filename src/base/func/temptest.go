package main

import "fmt"

type stu struct {
	name string
	age  int
}

func (s *stu) info(b string) {
	s.name = b
}

func main() {
	var s stu
	s.info("test")
	fmt.Println(s)
}
