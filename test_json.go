package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {

	p := Person{
		Name:  "zhangyu",
		Age:   23,
		Email: "aaaaaa@qq.com",
	}

	b, _ := json.Marshal(p)
	fmt.Printf("string(b): %v\n", string(b))
	s := `{"Name":"zhangyu","Age":23,"Email":"aaaaaa@qq.com"}`
	var m Person
	json.Unmarshal([]byte(s), &m)
	fmt.Printf("m: %v\n", m)
}
