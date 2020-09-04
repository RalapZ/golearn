package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	filename := "openfile.txt"
	FileR, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	err = errors.New("this is an define error")
	//fmt.Println(err.Error())
	if err != nil {
		//fmt.Println("test")
		//patherr,_:= err.(*os.PathError)
		//fmt.Println(patherr)

		fmt.Errorf("%s", err.Error())
		panic(err.Error())
		//fmt.Errorf("%s",err.Error())
	}

	//test,errz
	defer FileR.Close()
	//filebuf,_:=ioutil.ReadAll(FileR)
	//fmt.Println(string(filebuf))
}
