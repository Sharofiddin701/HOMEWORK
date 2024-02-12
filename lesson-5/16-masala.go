package main

import "fmt"

func main() {

	var a float32
	var k float32 = 1

	fmt.Print("a= ")
	fmt.Scan(&a)

	for i := 1; i <= 10; i++ {

		k = float32(i) * a

		fmt.Println(i, "kg konfet narxi:", k)
	}

}
