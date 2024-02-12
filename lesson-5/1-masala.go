package main

import "fmt"

func main() {

	var W int

	fmt.Print("a= ")
	fmt.Scan(&W)

	if W%3 == 0 {
		fmt.Println("FIZZ")
	} else if W%5 == 0 {
		fmt.Println("BUZZ")
	} else if W%3 == 0 || (W%5 == 0) {
		fmt.Println("FIZZBUZZ")
	}
}
