package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name": "murray", "age":25}`
	var p person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		panic("json unmarshal error")
	}
	fmt.Println(p)
}
