package main

import "fmt"

func main() {
	fmt.Println()
}

func thirdTask(numbers ...int) (positive, negative, even, odd []int) {
	for _, num := range numbers {
		if num > 0 {
			positive = append(positive, num)
		} else if num < 0 {
			negative = append(negative, num)
		}

		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return
}
