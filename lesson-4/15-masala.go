package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("a=")
	fmt.Scan(&a)
	fmt.Println("b=")
	fmt.Scan(&b)

	k := 2 * (a + b)

	fmt.Println("To'rtburchakning perimetri=", k)

}
