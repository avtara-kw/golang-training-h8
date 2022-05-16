package main

import "fmt"

func main(){
	fmt.Println("<isi pesan>")
	fmt.Println("Hello, world!")
	fmt.Println("Hello World", "Message", "From", "Go")

	fmt.Println("Avtara Khrisna")
	fmt.Println("Avtara", "Khrisna")

	fmt.Print("Avtara Khrisna\n")
	fmt.Print("Avtara", "Khrisna\n")

	//	komentar
	/*
		komentar
	*/

	// variable
	var name string = "avtara"
	var age int = 20

	fmt.Println("Ini adalah", name)
	fmt.Println("Umurnya", age)

	// assign value
	name = "joko"
	age = 30

	var address = "bandung"
	var postalCode = 53121

	fmt.Printf("%T %T\n", address, postalCode)

	// short declaration
	isValid := true
	fmt.Printf("%T\n", isValid)
	// multiple variable declarations
	var student1, student2, student3 string = "student1", "student2", "student3"
	var first, second int
	first, second = 1, 2

	fmt.Println(student1, student2, student3)

	// underscore variable
	_, _ = first, second

	// data types
	var positive uint8 = 20
	var negative int8 = -20

	fmt.Printf("%T %T\n", positive, negative)

	var decimalNumber float64 = 0.23
	fmt.Printf("%f\n", decimalNumber)
	fmt.Printf("%.3f\n", decimalNumber) 

	var condition bool = true
	fmt.Printf("is permitted? %t\n", condition)

	var message string = "ini string"
	fmt.Printf("is permitted? %s\n", message)

	var message1 string = `testing
	
lucu`
	fmt.Println(message1)

	// constant
	const PHI = 3.14
	fmt.Println(PHI)

	// Operators
	var value = (2 + 2) * 3
	fmt.Println(value)

	// Operators (Relational Operators)
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("First Condition:", firstCondition)
	fmt.Println("Second Condition:", secondCondition)
	fmt.Println("Third Condition:", thirdCondition)
	fmt.Println("Fourth Condition:", fourthCondition)

	var right, wrong = true, false
	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right\t(%t) \n", wrongAndRight)
	
	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right\t(%t) \n", wrongOrRight)
	
	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}