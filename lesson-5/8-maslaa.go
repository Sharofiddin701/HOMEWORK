package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)
	fmt.Print("c= ")
	fmt.Scan(&c)

	if a > 0 {
		a = 1
	} else {
		a = 0
	}

	if b > 0 {
		b = 1
	} else {
		b = 0
	}

	if c > 0 {
		c = 1
	} else {
		c = 0
	}
	fmt.Println("ta musbat son bor", a+b+c)

}
