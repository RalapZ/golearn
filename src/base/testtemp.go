package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("a.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	fmt.Println("hello ")
}
