package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string
}

func (s *student) SetName(name string) {
	s.name = name
}

func (s *student) GetName() string {
	return s.name
}

func main() {
	var s1 = &student{name: "Hacktiv8"}
	fmt.Println("name :", s1.name)

	reflectValue := reflect.ValueOf(s1)
	method := reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("koinworks"),
	})

	// s1.SetName("coinworks")
	fmt.Println("name :", s1.name)
}
