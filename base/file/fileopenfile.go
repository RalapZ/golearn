package main

import (
	"bufio"
	"os"
)

func main() {
	filename := "openfile.txt"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err.Error())
	}
	test := "test\n"
	defer file.Close()
	//file.Seek(0,2)
	//file.WriteString(test)
	fileWrite := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		test = test
		fileWrite.WriteString(test)
	}
	fileWrite.Flush()
}
