package main

import (
	"fmt"
)

func main() {
	var n, son int

	fmt.Println("n=")
	fmt.Scan(&n)

	slice := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Print("son=")
		fmt.Scan(&son)

		slice[i] = son
	}
	minnum := slice[0]
	for _, num := range slice {
		if num < minnum {
			minnum = num

		}
	}
	for i := 0; i < n; i++ {
		if slice[i] == minnum {
			fmt.Println("mindan oldin", i, "ta son bor")
		}
	}
	fmt.Println("minnum=", minnum)

}
