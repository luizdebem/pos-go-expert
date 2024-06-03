package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Array []int  `json:"array"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{"user": { "id": 123, "name": "John Doe", "active": true, "array": [1,2,3], "age": 25 } }`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	userJSON := value.Get("user").String()
	var user User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}
	fmt.Println(user.Name)
	// fmt.Printf("user name: %s\n", user.Get("name"))
	// fmt.Printf("user age: %s\n", user.Get("age"))

}
