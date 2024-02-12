package main

import "fmt"

func main() {

	var W int

	fmt.Print("a= ")
	fmt.Scan(&W)

	if W%2 == 0 {
		fmt.Println("Son juft", W*W)
	} else {
		fmt.Println("Son toq", W*W*W)
	}
}
