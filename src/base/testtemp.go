package main

import (
	"bytes"
	"io"
)

const debug = false

type a struct {
	t int
}

func main() {
	//io.Reader()
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
	//f1:=new(a)
	//f2:=make([]a,5)
	//fmt.Printf("%v",f1)
	//fmt.Printf("%v",f2)
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
