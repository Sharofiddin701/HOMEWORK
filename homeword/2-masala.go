package main

import "fmt"

func main() {

	var n int

	fmt.Print("n=")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {

		if i%2 == 0 {

			fmt.Println(i)

		}

	}

}
