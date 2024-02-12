package main

import "fmt"

func main() {

	var W int
	var s, t string

	fmt.Print("a= ")
	fmt.Scan(&W)
	fmt.Print("Ismni kiriting:")
	fmt.Scan(&s)
	fmt.Print("Familiyani kiriting:")
	fmt.Scan(&t)

	if W%2 == 0 {
		fmt.Println("Son juft :", t)
	} else {
		fmt.Println("Son toq : ", s)
	}
}
