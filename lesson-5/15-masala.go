package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)

	for i := b; i <= a; i-- {

		fmt.Println(i)

	}

}
