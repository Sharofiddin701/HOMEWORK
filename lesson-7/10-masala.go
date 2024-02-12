package main

import "fmt"

func kattason(a, b, c int) {

	var min, max int

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	if max > c {
		fmt.Print(" Katta son=", max)
	} else {
		fmt.Print(" Katta son=", c)
	}
	if min < c {
		fmt.Println(" Kichik son=", min)
	} else {
		fmt.Println(" Kichik son= ", c)
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

	kattason(a, b, c)

}
