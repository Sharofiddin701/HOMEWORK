package main

import "fmt"

func main() {

	var i, j, a int

	fmt.Print("Kiriting = ")
	fmt.Scanln(&a)

	for i = 1; i <= a; i++ {
		for j = 1; j <= a-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
