package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func main() {
	person := Person{Name: "polarisxu", Hobby: "Golang"}
	m, _ := json.Marshal(person)
	fmt.Printf("%s", m)
}
