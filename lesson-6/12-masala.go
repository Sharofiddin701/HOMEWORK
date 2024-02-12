package main

import "fmt"

func sonlar(a, b, c int) {
	if a+b == c {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}

func main() {

	var a, b, c int

	fmt.Print("A=")
	fmt.Scan(&a)
	fmt.Print("B=")
	fmt.Scan(&b)
	fmt.Print("C=")
	fmt.Scan(&c)

	sonlar(a, b, c)

}
