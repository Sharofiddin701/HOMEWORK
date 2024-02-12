package main

import "fmt"

func main() {

	var n int

	fmt.Print("1 dan farqli son kiriting !!! :")
	fmt.Scan(&n)

	tub := true

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			tub = false
		}
	}

	if tub {
		fmt.Printf("Tub son")
	} else {
		fmt.Printf("Tub son emas !!! \n")
	}
}
