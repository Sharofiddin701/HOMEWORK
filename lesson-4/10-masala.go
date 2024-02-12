package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Kattasi=", a)
	} else if a == b {
		fmt.Println("Sonlar teng")
	} else {
		fmt.Println("Kattasi=", b)
	}
}
