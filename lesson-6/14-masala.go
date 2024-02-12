package main

import "fmt"

func son(a float64, b int) {

	var k float64 = 1

	for i := 1; i <= b; i++ {

		k *= a

	}
	fmt.Println(k)

}
func main() {

	var a float64
	var b int

	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)

	son(a, b)

}
