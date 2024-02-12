package main

import "fmt"

func main() {

	for i := 10; i <= 99; i++ {

		tub := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				tub = false
			}
		}

		if tub {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println()
}
