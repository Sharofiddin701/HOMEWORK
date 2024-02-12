package main

import "fmt"

func main() {

	var a string

	fmt.Print("Kiriting :")
	fmt.Scan(&a)

	k := a

	for i := 0; i < len(k); i++ {
		fmt.Print(string(k[i]), " ")
	}
}
