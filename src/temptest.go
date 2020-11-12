package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	fmt.Println(w.(*os.File))      // success: f == os.Stdout
	fmt.Println(w.(*bytes.Buffer)) //
}
