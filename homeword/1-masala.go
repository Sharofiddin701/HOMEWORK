package main

import "fmt"

func main() {

	var a, b, c int

	for i := 100; i < 999; i++ {

		a = i / 100
		b = (i % 100) / 10
		c = (i % 100) % 10

		if (a == b) || (b == c) {
			fmt.Println("i=", i)
		}

	}

}
