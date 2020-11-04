package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

//使用File自带的Read
func read1(filename string) int {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 4096)
	var nbytes int
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		nbytes += n
	}
	return nbytes
}

//使用bufio
func read2(filename string) int {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 4096)
	var nbytes int
	rd := bufio.NewReader(fi)
	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		nbytes += n
	}
	return nbytes
}

//使用ioutil
func read3(filename string) int {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	nbytes := len(fd)
	return nbytes
}
func testfile1(filename string) {
	fmt.Printf("============test1 %s ===========\n", filename)
	start := time.Now()
	size1 := read1(filename)
	t1 := time.Now()
	fmt.Printf("Read 1 cost: %v, size: %d\n", t1.Sub(start), size1)
	size2 := read2(filename)
	t2 := time.Now()
	fmt.Printf("Read 2 cost: %v, size: %d\n", t2.Sub(t1), size2)
	size3 := read3(filename)
	t3 := time.Now()
	fmt.Printf("Read 3 cost: %v, size: %d\n", t3.Sub(t2), size3)
}

func main() {
	testfile1("small.txt")
	testfile1("midium.txt")
	testfile1("large.txt")
	// testfile2("small.txt")
	// testfile2("midium.txt")
	// testfile2("large.txt")
}
