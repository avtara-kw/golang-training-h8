package helpers

func Sum(num ...int) int {
	total := 0
	for _, n := range num {
		total += n
	}

	if total == 0 {
		total = -1
	}

	return total
}
