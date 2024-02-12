package main

import "fmt"

func main() {

	var juft, toq int

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	juft = slice1[2] * slice1[4] * slice1[6] * slice1[8]

	toq = slice1[1] + slice1[3] + slice1[5] + slice1[7] + slice1[9]

	fmt.Println(juft)
	fmt.Println(toq)
}
