package main

import "fmt"

func main() {

	
	fmt.Print(plusOne())
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {

		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
	}

	arr := []int{1}

	return append(arr, digits...)
}
