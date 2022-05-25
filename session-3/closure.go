package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{0, 5, 8, 10, 4, 2, 8, 2}
	var calculate = func(nums ...int) (total int) {
		for _, num := range nums {
			total += num
		}
		return
	}
	fmt.Println(calculate(nums...))

	students := []string{"Hacktiv8", "Koinworks", "NooBeeID"}

	student := findStudent(students)
	fmt.Println(student("noobee"))

	isOdd := func(i int) bool {
		return i%2 != 0
	}

	oddNumbers := findOddNumbers(nums, isOdd)

	fmt.Println(oddNumbers)
}

func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for k, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = k + 1
				break
			}

		}

		if student == "" {
			return fmt.Sprintf("Student dengan nama %s tidak ditemukan", s)
		}

		return fmt.Sprintf("student degnan nama %s ditemukan di posisi %d", s, position)
	}
}

func findOddNumbers(numbers []int, isOdd func(int) bool) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if isOdd(v) {
			totalOddNumbers += v
		}
	}

	return totalOddNumbers
}
