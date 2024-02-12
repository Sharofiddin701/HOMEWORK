package main

import (
	"fmt"
)

func main() {

	var a float64
	pi := 3.14

	fmt.Print("a=")
	fmt.Scan(&a)

	k := 2 * pi * a

	fmt.Printf("Uzunligi= %v\n", k)

}
