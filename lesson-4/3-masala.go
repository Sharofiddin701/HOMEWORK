package main

import (
	"fmt"
)

func main() {

	var a, b, c int
	
	fmt.Println("a=")
	fmt.Scan(&a)
	fmt.Println("b=")
	fmt.Scan(&b)
	fmt.Println("c=")
	fmt.Scan(&c)

	k := ((a+b > c) && (a+c > b) && (b+c > a))

	fmt.Println(k)

}
