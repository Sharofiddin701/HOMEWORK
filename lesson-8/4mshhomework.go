package main

import "fmt"

func main() {

	var n, kichik int

	fmt.Println("n=")
	fmt.Scan(&n)

	slice1 := make([]int, n)

	for i := 0; i < len(slice1)-1; i++ {

		if slice1[i] > slice1[i+1] {

			kichik = slice1[i]

		}
	}

	fmt.Print("Eng kichik elementi:", kichik)
}
