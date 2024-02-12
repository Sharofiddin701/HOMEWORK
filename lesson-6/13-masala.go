package main

import "fmt"

func toq(d int) {

	for i := 1; i <= d; i++ {

		if i%2 == 1 {

			fmt.Print(i, " ")

		}
	}

}
func main() {

	var d int

	fmt.Print("N=")
	fmt.Scan(&d)

	toq(d)

}
