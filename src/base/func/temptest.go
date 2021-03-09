package main

import (
	"fmt"
	"strings"
)

type intGen func() int

func (t intGen) Read(p []byte) (n int, err error) {
	test := t()
	s := fmt.Sprintf("%d\n", test)
	return strings.NewReader(s).Read(p)

}

func create() (fs [2]func()) {
	for i := 0; i < 2; i++ {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	return
}

func main() {
	fs := create()
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}
}
