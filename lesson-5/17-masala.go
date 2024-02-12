package main

import "fmt"

func main() {

	var a float32
	var k float32 = 1

	fmt.Print("a= ")
	fmt.Scan(&a)

	for i := 0.1; i <= 1; i += 0.1 {

		k = float32(i) * a

		fmt.Println(i, "kg konfet narxi:", k)
	}

}
