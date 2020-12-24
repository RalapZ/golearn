package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var (
	filenumber1      int
	fileexistnumber1 int
	filename         = "main.go"
	workernum        int
	workermaxnum     int
	workerdone       = make(chan bool)
	workernew        = make(chan bool)
	gonumchan        = make(chan int)
	gomatch          = make(chan string)
	pathstr          = make(chan string)
)

func worker() {
	for {
		select {
		case <-workerdone:
			workernum--
			if workernum == 0 {
				return
			}
		case <-workernew:
			workernum++
		case path := <-gomatch:
			go Searchchan(filename, path)
		}
	}
}

func Searchchan(filename, pathname string) {
	filelist, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range filelist {
		if file.IsDir() {
			if workernum < workermaxnum {
				gomatch <- pathname + "/" + file.Name()
			} else {
				Searchchan(filename, pathname+"/"+file.Name())
			}

		} else {
			if file.Name() == filename {
				fileexistnumber1++
			}
		}
	}
}

func main() {
	now := time.Now()
	pathname := "./"
	go Searchchan(filename, pathname)
	worker()
	fmt.Println(filenumber1, fileexistnumber1)
	fmt.Println(time.Since(now))
}
