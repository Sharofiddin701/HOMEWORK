package main

import "fmt"

func main() {

	var n int

	fmt.Print("So'zni kiriting:")
	fmt.Scan(&n)

	harf := make([]string, 0, n)

	for i := 0; i < n; i++ {

		temp := " "

		fmt.Scan(&temp)

		harf = append(harf, temp)
	}
	fmt.Println(harf)

}
