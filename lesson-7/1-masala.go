package main

import "fmt"

func kopaytma(a int) int {

	k := 1

	for i := 1; i <= a; i++ {

		k *= i
	}
	return k
}

func main() {
	var a int
	fmt.Print("Kiriting:")
	fmt.Scan(&a)

	fmt.Println(kopaytma(a))
}
