package main

import "fmt"

func array(x [10]int) int {
	var sum int

	sum = x[1] + x[3]

	return sum
}

func main() {
	var x [10]int
	x[0] = 1
	x[1] = -2
	x[2] = 3
	x[3] = -4
	x[4] = 5
	x[5] = -6
	x[6] = 7
	x[7] = -8
	x[8] = 9
	x[9] = -10

	fmt.Println(array(x))
}
