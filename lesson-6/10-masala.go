package main

import "fmt"

func son(n string) {

	for i := 1; i <= len(n); i++ {

		fmt.Print(n[i], " ")

	}

}
func main() {

	var n string

	fmt.Print("N=")
	fmt.Scan(&n)

	son(n)

}
