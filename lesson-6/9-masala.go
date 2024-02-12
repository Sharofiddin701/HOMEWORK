package main

import "fmt"

func juft(d int) {

	k := 1

	for i := 1; i <= d; i++ {

		k *= i

	}
	fmt.Println(k)

}
func main() {

	var d int

	fmt.Print("N=")
	fmt.Scan(&d)

	juft(d)

}
