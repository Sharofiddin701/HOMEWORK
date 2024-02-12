package main

import "fmt"

func main() {

	var i, j, a int

	fmt.Print("Kiriting = ")
	fmt.Scanln(&a)

	for i = a; i >= 1; i-- {
		for j = 0; j != (2*i - 1); j++ {
			fmt.Print(" ")
		}
		for k := 1; k >= a-i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
