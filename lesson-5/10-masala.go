package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Kattasi a ", a)
	} else {
		fmt.Println("Kattasi b ", b)
	}
}
