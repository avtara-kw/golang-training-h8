package main

import "fmt"

func main() {
	greet("Hacktiv8", "Hai")
	greet("Koinworks", "Hello")
	greet("Telkom", "Wow")

	total, err := calculate(10, 20, 50, 6, -8, 10)
	if err != nil {
		fmt.Println("error :", err.Error())
	}
	fmt.Println("total :", total)
}

func greet(name, text string) {
	fmt.Println(text, name)
}

func calculate(nums ...int) (total int, err error) {
	for key, num := range nums {
		if num < 0 {
			err = fmt.Errorf("data ke %d bernilai dibawah 0, yaitu :%d", key+1, num)
			return 0, err
		}
		total += num
	}
	return
}
