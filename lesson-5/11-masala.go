package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)

	if a < b {
		fmt.Println("1-da turgan son kichik ", a)
	} else {
		fmt.Println(" 2-da turgan son kichik", b)
	}
}
