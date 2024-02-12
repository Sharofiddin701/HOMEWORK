package main

import "fmt"

func son(a, b float64) {

	var k float64

	k = a / b

	fmt.Println(k)

}
func main() {

	var a, b float64

	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)

	son(a, b)

}
