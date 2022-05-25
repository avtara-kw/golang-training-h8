package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonString := `
		[
			{
				"name":"Hacktiv8",
				"age": 20
			},
			{
				"name":"Koinworks",
				"age": 22
			}
		]
	`

	// var result []Employee
	var result = []map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)

	stringJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stringJson))
}
