package main

import "fmt"

func kattason(a, b, c int) {

	var min, max, minn, maxx int

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	if max > c {
		fmt.Print(" Katta son=", max)
		maxx = max
	} else {
		fmt.Print(" Katta son=", c)
		maxx = c
	}
	if min < c {
		fmt.Println(" Kichik son=", min)
		minn = min
	} else {
		fmt.Println(" Kichik son= ", c)
		minn = c
	}
	for i := minn + 1; i < maxx; i++ {

		fmt.Print(i, " ")
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
