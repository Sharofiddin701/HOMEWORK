package main

import "fmt"

func son(x int) {

	a := 0

	for i := 1; i <= x; i++ {

		a += i

		fmt.Print(i, "+")

	}
	fmt.Print("=", a)
}

func main() {

	var x int

	fmt.Print("N=")
	fmt.Scan(&x)

	son(x)
}
