package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "test str-asd   myzone"
	re := regexp.MustCompile("(\\s)st.+?")
	k := re.FindAllStringSubmatch(a, 1)
	//i:=strings.TrimSpace(k[0][1])
	fmt.Println(k[0][0])
}
