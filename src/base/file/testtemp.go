package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()
	k := bufio.NewReader(file)
	buf := make([]byte, 4096)
	for {
		i, err := k.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		str := string(buf[:i])

		re := regexp.MustCompile("(?s)\\*[0-9][^*]+?\nLREM(?s).+?[^*]*")
		//fmt.Println(string(buf[:i]))
		t := re.FindAllStringSubmatch(str, -1)
		fmt.Println(t)
	}
}
