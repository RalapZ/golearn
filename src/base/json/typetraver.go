package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JsonTStruct(a string) {
	var p Person
	json.Unmarshal([]byte(a), &p)
	fmt.Printf("%T\n", p)
	fmt.Println(p)
}
func StructTJson(a Person) {
	data, _ := json.Marshal(a)
	fmt.Println(a)
	fmt.Println(string(data))
}

func MapConvertJson(M []map[string]interface{}) {
	k, _ := json.Marshal(M)
	fmt.Println(M)
	fmt.Println(string(k))
}

func main() {
	fmt.Println("json convert struct")
	a := `{"name":"ralap","age":"18"}`
	JsonTStruct(a)
	fmt.Println("struct convert json")
	p := Person{"ralapstuct", 28}
	StructTJson(p)
	fmt.Println("map convert json")
	test := []map[string]interface{}{}
	student1 := map[string]interface{}{"name": "ralapmap1", "age": 20, "sex": "male"}
	student2 := map[string]interface{}{"name": "ralapmap2", "age": 20, "sex": "male2"}
	test = append(test, student1)
	test = append(test, student2)
	MapConvertJson(test)

}
