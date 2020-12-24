package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var (
	filenumber      int
	fileexistnumber int
	workchan        = make(chan int)
)

func ReadAll(path string) {
	pathname, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range pathname {
		if file.IsDir() {
			go ReadAll(path + "/" + file.Name())
		} else {
			fmt.Println(file.Name())
			filenumber++
		}
	}
}

func Search(filename, pathname string) {
	filelist, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range filelist {
		if file.IsDir() {
			go Search(filename, pathname+"/"+file.Name())
		} else {
			if file.Name() == filename {
				fileexistnumber++
			}
		}
	}
}

func main() {
	now := time.Now()
	filename := "main.go"
	pathname := "./"
	ReadAll("./")
	Search(filename, pathname)
	fmt.Println(filenumber, fileexistnumber)
	fmt.Println(time.Since(now))
}
