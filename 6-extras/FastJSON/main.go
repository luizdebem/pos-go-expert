package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser

	jsonData := `{"id": 123, "name": "John Doe", "active": true, "array": [1,2,3]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\n", v.GetStringBytes("name"))

	fmt.Printf("ID: %d\n", v.GetInt64("id"))

	fmt.Printf("Active: %t\n", v.GetBool("active"))

	a := v.GetArray("array")
	for i, value := range a {
		fmt.Printf("Array[%d]: %d\n", i, value.GetInt64())
	}
}
