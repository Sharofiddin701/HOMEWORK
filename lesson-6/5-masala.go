package main

import "fmt"

func juft(d int) {

	for i := d; i >= 1; i-- {

		if i%2 == 0 {

			fmt.Println(i)

		}
	}

}
func main() {

	var d int

	fmt.Print("N=")
	fmt.Scan(&d)

	juft(d)

}
