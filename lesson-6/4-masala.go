package main

import "fmt"

func son(a int, b int) {

	if a < b {

		for i := a; i <= b; i++ {

			fmt.Print(i, " ")

		}

	} else if a > b {

		for i := a; i >= b; i-- {

			fmt.Print(i, " ")

		}

	}
}

func main() {

	var a, b int

	fmt.Print("A=")
	fmt.Scan(&a)
	fmt.Print("B=")
	fmt.Scan(&b)

	son(a, b)
}
