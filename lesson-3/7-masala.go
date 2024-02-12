package main

import "fmt"

func main() {

	var harf1, harf2, harf3 string

	fmt.Print("1-harfni kitiriting : ")
	fmt.Scan(&harf1)
	fmt.Print("2-harfni kitiriting : ")
	fmt.Scan(&harf2)
	fmt.Print("3-arfni kitiriting : ")
	fmt.Scan(&harf3)

	k := ((harf1 < harf2) && (harf2 < harf3 && (harf1 < harf3)))

	fmt.Printf("Unli: %v\n", k)
}
