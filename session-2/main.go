package main

import (
	"fmt"
)

func main(){
	var currentYear = 2021

	if age := currentYear - 1997; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	}else{
		fmt.Println("Kamu boleh membuat kartu sim")
	}

	var score = 4

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Good")
	default:
		fmt.Println("Not bad")
	}
	
	switch {
	case score == 8:
		fmt.Println("Perfect")
	case score < 8 && score > 3:
		fmt.Println("Not bad")
		fallthrough
	default:
		fmt.Println("study harder")
		fmt.Println("you need to learn more")
	}

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")
		}
	} else {
		if score == 5 {
			fmt.Println("Not Bad")
		} else if score == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if score == 0 {
				fmt.Println("Try harder")
			}
		}
	}

	// looping
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	var i = 0

	for i < 3 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 3 {
			break
		}
	}

	max := 50
	for i := 2; i <= max; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	// array
	var numbers = [4]int{1,2,3,4}
	fmt.Printf("%#v\n", numbers)
	numbers[2] = 99
	fmt.Printf("%#v\n", numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %d, val: %d\n", i, numbers[i])
	}

	balances := [2][3]int{{5,6,7},{8,9,10}}
	for _, arr := range balances {
		for _, val := range arr {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

	// slice
	var fruits = []string{"apple", "banana", "mango", "durian", "pineapple"}

	fmt.Printf("%#v\n", fruits[1:4])
	fmt.Printf("%#v\n", fruits[:3])
	fmt.Printf("%#v\n", fruits[1:])
	fmt.Printf("%#v\n", fruits[:])

	fruits = append(fruits[:2], "rambutan")
	fmt.Printf("%#v\n", fruits[:])

	// map
	var person map[string]string
	person = map[string]string{}

	person["name"] = "avtara"
	person["age"] = "20"
	fmt.Println(person["name"])

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
	delete(person, "age")
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	if value, ok := person["age"]; ok {
		fmt.Println("ada value", value)
	} else {
		fmt.Println("tidak ada value")
	}

	var people = []map[string]string{
		{"name": "avtara", "age": "20"},
		{"name": "joko", "age": "30"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name %s, age: %s\n", i, person["name"], person["age"])
	}

	// aliases
	type second = uint 
	var hour second = 3600
	fmt.Printf("%T", hour)
}